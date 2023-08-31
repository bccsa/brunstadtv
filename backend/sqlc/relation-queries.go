package sqlc

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// GetSongs returns songs for the specified ids
func (q *Queries) GetSongs(ctx context.Context, ids []uuid.UUID) ([]common.Song, error) {
	rows, err := q.getSongs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i Song, _ int) common.Song {
		title := toLocaleString(nil, i.Title)
		return common.Song{
			ID:    i.ID,
			Title: title,
		}
	}), nil
}

// GetPersons returns persons for the specified ids
func (q *Queries) GetPersons(ctx context.Context, ids []uuid.UUID) ([]common.Person, error) {
	rows, err := q.getPersons(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(i Person, _ int) common.Person {
		return common.Person{
			ID:   i.ID,
			Name: i.Name,
		}
	}), nil
}
