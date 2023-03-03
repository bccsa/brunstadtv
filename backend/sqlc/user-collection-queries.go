package sqlc

import (
	"context"
	"encoding/json"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/loaders"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// GetUserCollections returns user-collections by ids
func (q *Queries) GetUserCollections(ctx context.Context, ids []uuid.UUID) ([]common.UserCollection, error) {
	rows, err := q.getUserCollections(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i getUserCollectionsRow, _ int) common.UserCollection {
		var metadata common.UserCollectionMetadata
		_ = json.Unmarshal(i.Metadata.RawMessage, &metadata)
		return common.UserCollection{
			ID:        i.ID,
			ProfileID: i.ProfileID,
			Title:     i.Title,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
			Metadata:  metadata,
		}
	}), nil
}

// GetUserCollectionEntries returns entries by id
func (q *Queries) GetUserCollectionEntries(ctx context.Context, ids []uuid.UUID) ([]common.UserCollectionEntry, error) {
	rows, err := q.getUserCollectionEntries(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i getUserCollectionEntriesRow, _ int) common.UserCollectionEntry {
		return common.UserCollectionEntry{
			ID:     i.ID,
			Type:   i.Type,
			ItemID: i.ItemID,
			Sort:   int(i.Sort),
		}
	}), nil
}

// GetUserCollectionIDsForProfileIDs returns collection ids for profiles
func (q *Queries) GetUserCollectionIDsForProfileIDs(ctx context.Context, profileIDs []uuid.UUID) ([]loaders.Relation[uuid.UUID, uuid.UUID], error) {
	rows, err := q.getUserCollectionIDsForProfileIDs(ctx, profileIDs)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i getUserCollectionIDsForProfileIDsRow, _ int) loaders.Relation[uuid.UUID, uuid.UUID] {
		return relation[uuid.UUID, uuid.UUID](i)
	}), nil
}

// GetUserCollectionEntryIDsForUserCollectionIDs returns entry ids for collection ids
func (q *Queries) GetUserCollectionEntryIDsForUserCollectionIDs(ctx context.Context, collectionIDs []uuid.UUID) ([]loaders.Relation[uuid.UUID, uuid.UUID], error) {
	rows, err := q.getUserCollectionEntryIDsForUserCollectionIDs(ctx, collectionIDs)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i getUserCollectionEntryIDsForUserCollectionIDsRow, _ int) loaders.Relation[uuid.UUID, uuid.UUID] {
		return relation[uuid.UUID, uuid.UUID](i)
	}), nil
}
