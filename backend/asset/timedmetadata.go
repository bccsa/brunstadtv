package asset

import (
	"context"

	"github.com/bcc-code/bcc-media-platform/backend/sqlc"
	"github.com/bcc-code/bcc-media-platform/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"gopkg.in/guregu/null.v4"

	"github.com/bcc-code/bcc-media-platform/backend/common"

	"github.com/ansel1/merry/v2"
	"github.com/google/uuid"
)

type IngestTimedMetadataParams struct {
	VXID     string
	JSONPath string
}

// Ingest timedmetadata from a JSON file based on the vxID
func IngestTimedMetadata(ctx context.Context, services externalServices, config config, params IngestTimedMetadataParams) error {
	queries := services.GetQueries()
	s3client := services.GetS3Client()
	db := services.GetDatabase()
	tx, err := db.Begin()
	if err != nil {
		return merry.Wrap(err)
	}
	defer tx.Rollback()
	qtx := queries.WithTx(tx)

	assetIDs, err := qtx.AssetIDsByMediabankenID(ctx, params.VXID)
	if err != nil {
		return merry.Wrap(err)
	}

	if len(assetIDs) == 0 {
		return merry.New("no assets found", merry.WithUserMessage("no assets found for VXID: "+params.VXID))
	}

	for _, assetID := range assetIDs {
		err = qtx.ClearAssetTimedMetadata(ctx, null.IntFrom(int64(assetID)))
		if err != nil {
			return merry.Wrap(err)
		}
	}

	newestAssetID := assetIDs[0]

	var timedMetadatas []TimedMetadata
	err = readJSONFromS3(ctx, s3client, config.GetIngestBucket(), params.JSONPath, &timedMetadatas)
	if err != nil {
		return merry.Wrap(err)
	}

	for _, inputTm := range timedMetadatas {
		imageFileID, err := generateImageForAssetAtTime(ctx, services, config, newestAssetID, inputTm.Timestamp)
		if err != nil {
			return merry.Wrap(err)
		}
		err = ingestOneTimedMetadata(ctx, qtx, inputTm, assetIDs, imageFileID)
		if err != nil {
			return merry.Wrap(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return merry.Wrap(err)
	}

	return nil
}

// ingestOneTimedMetadata creates one timed metadata with contributions, song, etc.
func ingestOneTimedMetadata(ctx context.Context, queries *sqlc.Queries, inputTm TimedMetadata, assetIDs []int32, imageID *string) error {
	t := common.ContentTypes.Parse(inputTm.ContentType)
	if t == nil {
		log.L.Warn().Msg("Skipping. Unknown chapter type: " + inputTm.ContentType)
		return nil
	}
	realTm := sqlc.InsertTimedMetadataParams{
		ContentType: null.StringFrom(t.Value),
		Title:       inputTm.Title,
		Highlight:   inputTm.Highlight,
		Description: inputTm.Description,
		Status:      string(common.StatusPublished),
		Label:       inputTm.Label,
		Type:        "chapter",
		Seconds:     float32(inputTm.Timestamp),
	}

	var personIDs []uuid.UUID
	personIDs, err := getOrInsertPersonIDs(ctx, queries, inputTm.Persons)
	if err != nil {
		return merry.Wrap(err)
	}

	if inputTm.SongCollection != "" && inputTm.SongNumber != "" {
		songID, err := getOrInsertSongID(ctx, queries, inputTm.SongCollection, inputTm.SongNumber)
		if err != nil {
			return merry.Wrap(err)
		}
		realTm.SongID = uuid.NullUUID{
			Valid: true,
			UUID:  songID,
		}
	}

	for _, assetID := range assetIDs {
		realTm.ID = uuid.New()
		realTm.AssetID = null.IntFrom(int64(assetID))
		tmID, err := queries.InsertTimedMetadata(ctx, realTm)
		if err != nil {
			return merry.Wrap(err)
		}
		for _, p := range personIDs {
			err = queries.InsertContribution(ctx, sqlc.InsertContributionParams{
				PersonID:        p,
				Type:            MapContributionTypeFromContentType(*t).Value,
				TimedmetadataID: uuid.NullUUID{UUID: tmID, Valid: true},
			})

			if err != nil {
				return merry.Wrap(err)
			}
		}
		if imageID != nil {
			styledId, err := queries.InsertStyledImage(ctx, sqlc.InsertStyledImageParams{
				Language: (*utils.FallbackLanguages())[0],
				Style:    common.ImageStyleDefault,
				File:     uuid.MustParse(*imageID),
			})
			if err != nil {
				return merry.Wrap(err)
			}

			_, err = queries.InsertTimedMetadataStyledImage(ctx, sqlc.InsertTimedMetadataStyledImageParams{
				TimedMetadataID: uuid.NullUUID{UUID: tmID, Valid: true},
				StyledImageID:   uuid.NullUUID{UUID: styledId, Valid: true},
			})
			if err != nil {
				return merry.Wrap(err)
			}
		}
	}
	return nil
}

// MapContributionTypeFromContentType guesses the contribution type based on the chapter type
func MapContributionTypeFromContentType(t common.ContentType) common.ContributionType {
	switch t {
	case common.ContentTypeInterview, common.ContentTypeSpeech, common.ContentTypeTestimony, common.ContentTypeTheme:
		return common.ContributionTypeSpeaker
	case common.ContentTypeSingAlong, common.ContentTypeSong:
		return common.ContributionTypeSinger
	}

	return common.ContributionTypeUnknown
}
