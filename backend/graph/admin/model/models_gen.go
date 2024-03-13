// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CollectionItem struct {
	Collection Collection `json:"collection"`
	ID         string     `json:"id"`
	Title      string     `json:"title"`
}

type Episodes struct {
	ImportTimedMetadata bool `json:"importTimedMetadata"`
}

type MediaItems struct {
	ImportTimedMetadata bool `json:"importTimedMetadata"`
}

type Preview struct {
	Collection *PreviewCollection `json:"collection"`
	Asset      *PreviewAsset      `json:"asset"`
}

type PreviewAsset struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type PreviewCollection struct {
	Items []*CollectionItem `json:"items"`
}

type ProgressByOrg struct {
	Name     string  `json:"name"`
	Progress float64 `json:"progress"`
}

type Statistics struct {
	LessonProgressGroupedByOrg []*ProgressByOrg `json:"lessonProgressGroupedByOrg"`
}

type Collection string

const (
	CollectionShows    Collection = "shows"
	CollectionSeasons  Collection = "seasons"
	CollectionEpisodes Collection = "episodes"
	CollectionGames    Collection = "games"
	CollectionShorts   Collection = "shorts"
)

var AllCollection = []Collection{
	CollectionShows,
	CollectionSeasons,
	CollectionEpisodes,
	CollectionGames,
	CollectionShorts,
}

func (e Collection) IsValid() bool {
	switch e {
	case CollectionShows, CollectionSeasons, CollectionEpisodes, CollectionGames, CollectionShorts:
		return true
	}
	return false
}

func (e Collection) String() string {
	return string(e)
}

func (e *Collection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Collection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Collection", str)
	}
	return nil
}

func (e Collection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
