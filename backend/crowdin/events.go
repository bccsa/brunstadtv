package crowdin

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/directus"
	"github.com/bcc-code/brunstadtv/backend/events"
	"github.com/bcc-code/mediabank-bridge/log"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
)

var (
	// ErrCollectionNotSupported error for collection not yet supported
	ErrCollectionNotSupported = merry.Sentinel("collection not supported")
	// ErrEventNotSupported error for event key not yet supported
	ErrEventNotSupported = merry.Sentinel("event not supported")
)

type services interface {
	GetCrowdinClient() *Client
	GetDirectusClient() *resty.Client
}

// HandleEvent for events from PubSub (or other CloudEvent source)
func HandleEvent(ctx context.Context, services services, event cloudevents.Event) (err error) {
	client := services.GetCrowdinClient()
	switch event.Type() {
	case events.TypeTranslationsSync:
		handler := directus.NewHandler(services.GetDirectusClient())
		return client.Sync(ctx, handler)
	}
	return merry.Wrap(ErrEventNotSupported)
}

func toTranslationSources[t TranslationSource](items []t) []TranslationSource {
	return lo.Map(items, func(i t, _ int) TranslationSource {
		return i
	})
}

func getStatusForItem(ctx context.Context, d *directus.Handler, collection string, id int) (status common.Status, err error) {
	var i hasStatus
	switch collection {
	case "shows":
		i, err = d.GetShow(ctx, id)
	case "seasons":
		i, err = d.GetSeason(ctx, id)
	case "episodes":
		i, err = d.GetEpisode(ctx, id)
	}
	if err == nil {
		status = i.GetStatus()
	}
	return
}

func getTranslationsForItem(ctx context.Context, d *directus.Handler, collection string, id int, language string) ([]TranslationSource, error) {
	switch collection {
	case "shows":
		ts, err := d.ListShowTranslations(ctx, language, false, id)
		if err != nil {
			return nil, err
		}
		return toTranslationSources(ts), nil
	case "seasons":
		ts, err := d.ListSeasonTranslations(ctx, language, false, id)
		if err != nil {
			return nil, err
		}
		return toTranslationSources(ts), nil
	case "episodes":
		ts, err := d.ListEpisodeTranslations(ctx, language, false, id)
		if err != nil {
			return nil, err
		}
		return toTranslationSources(ts), nil
	}
	return nil, merry.Wrap(ErrCollectionNotSupported)
}

// HandleModelUpdate for triggering actions on object change
func (client *Client) HandleModelUpdate(ctx context.Context, collection string, id int) error {
	if status, err := getStatusForItem(ctx, client.du, collection, id); err != nil || status != "published" {
		// Return error, else just ignore if not published
		return err
	}
	translations, err := getTranslationsForItem(ctx, client.du, collection, id, "")
	if err != nil {
		return err
	}
	if len(translations) == 0 {
		return nil
	}
	return client.SaveTranslations(translations)
}

// HandleModelDelete for triggering actions to handle deletion events
func (client *Client) HandleModelDelete(_ context.Context, collection string, id int) error {
	log.L.Debug().Str("collection", collection).Int("id", id).Msg("deleting translations - not implemented")
	// TODO: implement deletion of translations
	return nil
}
