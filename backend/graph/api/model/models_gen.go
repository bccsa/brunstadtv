// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CalendarEntry interface {
	IsCalendarEntry()
	GetID() string
	GetEvent() *Event
	GetTitle() string
	GetDescription() string
	GetStart() string
	GetEnd() string
}

type CollectionItem interface {
	IsCollectionItem()
	GetID() string
	GetSort() int
	GetTitle() string
	GetImageURL() *string
	GetImages() []*Image
}

type GridSection interface {
	IsSection()
	IsItemSection()
	IsGridSection()
	GetID() string
	GetMetadata() *ItemSectionMetadata
	GetTitle() *string
	GetSize() GridSectionSize
	GetItems() *SectionItemPagination
}

type ItemSection interface {
	IsSection()
	IsItemSection()
	GetID() string
	GetMetadata() *ItemSectionMetadata
	GetTitle() *string
	GetItems() *SectionItemPagination
}

type Pagination interface {
	IsPagination()
	GetTotal() int
	GetFirst() int
	GetOffset() int
}

type SearchResultItem interface {
	IsSearchResultItem()
	GetID() string
	GetLegacyID() *string
	GetCollection() string
	GetTitle() string
	GetHeader() *string
	GetDescription() *string
	GetHighlight() *string
	GetImage() *string
	GetURL() string
}

type Section interface {
	IsSection()
	GetID() string
	GetTitle() *string
}

type SectionItemType interface {
	IsSectionItemType()
}

type Application struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	ClientVersion string `json:"clientVersion"`
	Page          *Page  `json:"page"`
	SearchPage    *Page  `json:"searchPage"`
}

type Calendar struct {
	Period *CalendarPeriod `json:"period"`
	Day    *CalendarDay    `json:"day"`
}

type CalendarDay struct {
	Events  []*Event        `json:"events"`
	Entries []CalendarEntry `json:"entries"`
}

type CalendarPeriod struct {
	ActiveDays []string `json:"activeDays"`
	Events     []*Event `json:"events"`
}

type Chapter struct {
	ID    string `json:"id"`
	Start int    `json:"start"`
	Title string `json:"title"`
}

type Collection struct {
	ID    string                    `json:"id"`
	Items *CollectionItemPagination `json:"items"`
}

type CollectionItemPagination struct {
	Total  int              `json:"total"`
	First  int              `json:"first"`
	Offset int              `json:"offset"`
	Items  []CollectionItem `json:"items"`
}

func (CollectionItemPagination) IsPagination()       {}
func (this CollectionItemPagination) GetTotal() int  { return this.Total }
func (this CollectionItemPagination) GetFirst() int  { return this.First }
func (this CollectionItemPagination) GetOffset() int { return this.Offset }

type Config struct {
	Global *GlobalConfig `json:"global"`
}

type DefaultGridSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     GridSectionSize        `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (DefaultGridSection) IsSection()             {}
func (this DefaultGridSection) GetID() string     { return this.ID }
func (this DefaultGridSection) GetTitle() *string { return this.Title }

func (DefaultGridSection) IsItemSection() {}

func (this DefaultGridSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this DefaultGridSection) GetItems() *SectionItemPagination { return this.Items }

func (DefaultGridSection) IsGridSection() {}

func (this DefaultGridSection) GetSize() GridSectionSize { return this.Size }

type DefaultSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     SectionSize            `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (DefaultSection) IsSection()             {}
func (this DefaultSection) GetID() string     { return this.ID }
func (this DefaultSection) GetTitle() *string { return this.Title }

func (DefaultSection) IsItemSection() {}

func (this DefaultSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this DefaultSection) GetItems() *SectionItemPagination { return this.Items }

type Device struct {
	Token     string `json:"token"`
	UpdatedAt string `json:"updatedAt"`
}

type Episode struct {
	ID                string                 `json:"id"`
	Type              EpisodeType            `json:"type"`
	LegacyID          *string                `json:"legacyID"`
	LegacyProgramID   *string                `json:"legacyProgramID"`
	PublishDate       string                 `json:"publishDate"`
	AvailableFrom     string                 `json:"availableFrom"`
	AvailableTo       string                 `json:"availableTo"`
	AgeRating         string                 `json:"ageRating"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description"`
	ExtraDescription  string                 `json:"extraDescription"`
	Image             *string                `json:"image"`
	ImageURL          *string                `json:"imageUrl"`
	ProductionDate    *string                `json:"productionDate"`
	Streams           []*Stream              `json:"streams"`
	Files             []*File                `json:"files"`
	Chapters          []*Chapter             `json:"chapters"`
	Season            *Season                `json:"season"`
	Duration          int                    `json:"duration"`
	Progress          *int                   `json:"progress"`
	AudioLanguages    []Language             `json:"audioLanguages"`
	SubtitleLanguages []Language             `json:"subtitleLanguages"`
	Context           *SectionItemPagination `json:"context"`
	RelatedItems      *SectionItemPagination `json:"relatedItems"`
	Images            []*Image               `json:"images"`
	Number            *int                   `json:"number"`
}

func (Episode) IsSectionItemType() {}

type EpisodeCalendarEntry struct {
	ID          string   `json:"id"`
	Event       *Event   `json:"event"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Start       string   `json:"start"`
	End         string   `json:"end"`
	Episode     *Episode `json:"episode"`
}

func (EpisodeCalendarEntry) IsCalendarEntry()            {}
func (this EpisodeCalendarEntry) GetID() string          { return this.ID }
func (this EpisodeCalendarEntry) GetEvent() *Event       { return this.Event }
func (this EpisodeCalendarEntry) GetTitle() string       { return this.Title }
func (this EpisodeCalendarEntry) GetDescription() string { return this.Description }
func (this EpisodeCalendarEntry) GetStart() string       { return this.Start }
func (this EpisodeCalendarEntry) GetEnd() string         { return this.End }

type EpisodeContext struct {
	CollectionID *string `json:"collectionId"`
}

type EpisodeItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Episode  *Episode `json:"episode"`
}

func (EpisodeItem) IsCollectionItem()         {}
func (this EpisodeItem) GetID() string        { return this.ID }
func (this EpisodeItem) GetSort() int         { return this.Sort }
func (this EpisodeItem) GetTitle() string     { return this.Title }
func (this EpisodeItem) GetImageURL() *string { return this.ImageURL }
func (this EpisodeItem) GetImages() []*Image {
	if this.Images == nil {
		return nil
	}
	interfaceSlice := make([]*Image, 0, len(this.Images))
	for _, concrete := range this.Images {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type EpisodePagination struct {
	Total  int        `json:"total"`
	First  int        `json:"first"`
	Offset int        `json:"offset"`
	Items  []*Episode `json:"items"`
}

func (EpisodePagination) IsPagination()       {}
func (this EpisodePagination) GetTotal() int  { return this.Total }
func (this EpisodePagination) GetFirst() int  { return this.First }
func (this EpisodePagination) GetOffset() int { return this.Offset }

type EpisodeSearchItem struct {
	ID              string  `json:"id"`
	LegacyID        *string `json:"legacyID"`
	LegacyProgramID *string `json:"legacyProgramID"`
	Duration        int     `json:"duration"`
	AgeRating       string  `json:"ageRating"`
	Collection      string  `json:"collection"`
	Title           string  `json:"title"`
	Header          *string `json:"header"`
	Description     *string `json:"description"`
	Highlight       *string `json:"highlight"`
	Image           *string `json:"image"`
	URL             string  `json:"url"`
	ShowID          *string `json:"showId"`
	ShowTitle       *string `json:"showTitle"`
	Show            *Show   `json:"show"`
	SeasonID        *string `json:"seasonId"`
	SeasonTitle     *string `json:"seasonTitle"`
	Season          *Season `json:"season"`
}

func (EpisodeSearchItem) IsSearchResultItem()          {}
func (this EpisodeSearchItem) GetID() string           { return this.ID }
func (this EpisodeSearchItem) GetLegacyID() *string    { return this.LegacyID }
func (this EpisodeSearchItem) GetCollection() string   { return this.Collection }
func (this EpisodeSearchItem) GetTitle() string        { return this.Title }
func (this EpisodeSearchItem) GetHeader() *string      { return this.Header }
func (this EpisodeSearchItem) GetDescription() *string { return this.Description }
func (this EpisodeSearchItem) GetHighlight() *string   { return this.Highlight }
func (this EpisodeSearchItem) GetImage() *string       { return this.Image }
func (this EpisodeSearchItem) GetURL() string          { return this.URL }

type Event struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
	Image string `json:"image"`
}

type Export struct {
	DbVersion string `json:"dbVersion"`
	URL       string `json:"url"`
}

type Faq struct {
	Categories *FAQCategoryPagination `json:"categories"`
	Category   *FAQCategory           `json:"category"`
	Question   *Question              `json:"question"`
}

type FAQCategory struct {
	ID        string              `json:"id"`
	Title     string              `json:"title"`
	Questions *QuestionPagination `json:"questions"`
}

type FAQCategoryPagination struct {
	Total  int            `json:"total"`
	First  int            `json:"first"`
	Offset int            `json:"offset"`
	Items  []*FAQCategory `json:"items"`
}

func (FAQCategoryPagination) IsPagination()       {}
func (this FAQCategoryPagination) GetTotal() int  { return this.Total }
func (this FAQCategoryPagination) GetFirst() int  { return this.First }
func (this FAQCategoryPagination) GetOffset() int { return this.Offset }

type FeaturedSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     SectionSize            `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (FeaturedSection) IsSection()             {}
func (this FeaturedSection) GetID() string     { return this.ID }
func (this FeaturedSection) GetTitle() *string { return this.Title }

func (FeaturedSection) IsItemSection() {}

func (this FeaturedSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this FeaturedSection) GetItems() *SectionItemPagination { return this.Items }

type File struct {
	ID               string    `json:"id"`
	URL              string    `json:"url"`
	AudioLanguage    Language  `json:"audioLanguage"`
	SubtitleLanguage *Language `json:"subtitleLanguage"`
	Size             *int      `json:"size"`
	FileName         string    `json:"fileName"`
	MimeType         string    `json:"mimeType"`
}

type GlobalConfig struct {
	LiveOnline  bool `json:"liveOnline"`
	NpawEnabled bool `json:"npawEnabled"`
}

type IconGridSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     GridSectionSize        `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (IconGridSection) IsSection()             {}
func (this IconGridSection) GetID() string     { return this.ID }
func (this IconGridSection) GetTitle() *string { return this.Title }

func (IconGridSection) IsItemSection() {}

func (this IconGridSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this IconGridSection) GetItems() *SectionItemPagination { return this.Items }

func (IconGridSection) IsGridSection() {}

func (this IconGridSection) GetSize() GridSectionSize { return this.Size }

type IconSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Items    *SectionItemPagination `json:"items"`
}

func (IconSection) IsSection()             {}
func (this IconSection) GetID() string     { return this.ID }
func (this IconSection) GetTitle() *string { return this.Title }

func (IconSection) IsItemSection() {}

func (this IconSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this IconSection) GetItems() *SectionItemPagination { return this.Items }

type Image struct {
	Style string `json:"style"`
	URL   string `json:"url"`
}

type ItemSectionMetadata struct {
	ContinueWatching bool   `json:"continueWatching"`
	SecondaryTitles  bool   `json:"secondaryTitles"`
	CollectionID     string `json:"collectionId"`
}

type LabelSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Items    *SectionItemPagination `json:"items"`
}

func (LabelSection) IsSection()             {}
func (this LabelSection) GetID() string     { return this.ID }
func (this LabelSection) GetTitle() *string { return this.Title }

func (LabelSection) IsItemSection() {}

func (this LabelSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this LabelSection) GetItems() *SectionItemPagination { return this.Items }

type LegacyIDLookup struct {
	ID string `json:"id"`
}

type LegacyIDLookupOptions struct {
	EpisodeID *int `json:"episodeID"`
	ProgramID *int `json:"programID"`
}

type Link struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (Link) IsSectionItemType() {}

type ListSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     SectionSize            `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (ListSection) IsSection()             {}
func (this ListSection) GetID() string     { return this.ID }
func (this ListSection) GetTitle() *string { return this.Title }

func (ListSection) IsItemSection() {}

func (this ListSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this ListSection) GetItems() *SectionItemPagination { return this.Items }

type Message struct {
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Style   *MessageStyle `json:"style"`
}

type MessageSection struct {
	ID       string               `json:"id"`
	Metadata *ItemSectionMetadata `json:"metadata"`
	Title    *string              `json:"title"`
	Messages []*Message           `json:"messages"`
}

func (MessageSection) IsSection()             {}
func (this MessageSection) GetID() string     { return this.ID }
func (this MessageSection) GetTitle() *string { return this.Title }

type MessageStyle struct {
	Text       string `json:"text"`
	Background string `json:"background"`
	Border     string `json:"border"`
}

type Page struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       string             `json:"title"`
	Description *string            `json:"description"`
	Image       *string            `json:"image"`
	Images      []*Image           `json:"images"`
	Sections    *SectionPagination `json:"sections"`
}

func (Page) IsSectionItemType() {}

type PageItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Page     *Page    `json:"page"`
}

func (PageItem) IsCollectionItem()         {}
func (this PageItem) GetID() string        { return this.ID }
func (this PageItem) GetSort() int         { return this.Sort }
func (this PageItem) GetTitle() string     { return this.Title }
func (this PageItem) GetImageURL() *string { return this.ImageURL }
func (this PageItem) GetImages() []*Image {
	if this.Images == nil {
		return nil
	}
	interfaceSlice := make([]*Image, 0, len(this.Images))
	for _, concrete := range this.Images {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type PosterGridSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     GridSectionSize        `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (PosterGridSection) IsSection()             {}
func (this PosterGridSection) GetID() string     { return this.ID }
func (this PosterGridSection) GetTitle() *string { return this.Title }

func (PosterGridSection) IsItemSection() {}

func (this PosterGridSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this PosterGridSection) GetItems() *SectionItemPagination { return this.Items }

func (PosterGridSection) IsGridSection() {}

func (this PosterGridSection) GetSize() GridSectionSize { return this.Size }

type PosterSection struct {
	ID       string                 `json:"id"`
	Metadata *ItemSectionMetadata   `json:"metadata"`
	Title    *string                `json:"title"`
	Size     SectionSize            `json:"size"`
	Items    *SectionItemPagination `json:"items"`
}

func (PosterSection) IsSection()             {}
func (this PosterSection) GetID() string     { return this.ID }
func (this PosterSection) GetTitle() *string { return this.Title }

func (PosterSection) IsItemSection() {}

func (this PosterSection) GetMetadata() *ItemSectionMetadata { return this.Metadata }

func (this PosterSection) GetItems() *SectionItemPagination { return this.Items }

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	ID       string       `json:"id"`
	Category *FAQCategory `json:"category"`
	Question string       `json:"question"`
	Answer   string       `json:"answer"`
}

type QuestionPagination struct {
	Total  int         `json:"total"`
	First  int         `json:"first"`
	Offset int         `json:"offset"`
	Items  []*Question `json:"items"`
}

func (QuestionPagination) IsPagination()       {}
func (this QuestionPagination) GetTotal() int  { return this.Total }
func (this QuestionPagination) GetFirst() int  { return this.First }
func (this QuestionPagination) GetOffset() int { return this.Offset }

type RedirectLink struct {
	URL string `json:"url"`
}

type RedirectParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SearchResult struct {
	Hits   int                `json:"hits"`
	Page   int                `json:"page"`
	Result []SearchResultItem `json:"result"`
}

type Season struct {
	ID          string             `json:"id"`
	LegacyID    *string            `json:"legacyID"`
	AgeRating   string             `json:"ageRating"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Image       *string            `json:"image"`
	ImageURL    *string            `json:"imageUrl"`
	Images      []*Image           `json:"images"`
	Number      int                `json:"number"`
	Show        *Show              `json:"show"`
	Episodes    *EpisodePagination `json:"episodes"`
}

func (Season) IsSectionItemType() {}

type SeasonCalendarEntry struct {
	ID          string  `json:"id"`
	Event       *Event  `json:"event"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Start       string  `json:"start"`
	End         string  `json:"end"`
	Season      *Season `json:"season"`
}

func (SeasonCalendarEntry) IsCalendarEntry()            {}
func (this SeasonCalendarEntry) GetID() string          { return this.ID }
func (this SeasonCalendarEntry) GetEvent() *Event       { return this.Event }
func (this SeasonCalendarEntry) GetTitle() string       { return this.Title }
func (this SeasonCalendarEntry) GetDescription() string { return this.Description }
func (this SeasonCalendarEntry) GetStart() string       { return this.Start }
func (this SeasonCalendarEntry) GetEnd() string         { return this.End }

type SeasonItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Season   *Season  `json:"season"`
}

func (SeasonItem) IsCollectionItem()         {}
func (this SeasonItem) GetID() string        { return this.ID }
func (this SeasonItem) GetSort() int         { return this.Sort }
func (this SeasonItem) GetTitle() string     { return this.Title }
func (this SeasonItem) GetImageURL() *string { return this.ImageURL }
func (this SeasonItem) GetImages() []*Image {
	if this.Images == nil {
		return nil
	}
	interfaceSlice := make([]*Image, 0, len(this.Images))
	for _, concrete := range this.Images {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type SeasonPagination struct {
	Total  int       `json:"total"`
	First  int       `json:"first"`
	Offset int       `json:"offset"`
	Items  []*Season `json:"items"`
}

func (SeasonPagination) IsPagination()       {}
func (this SeasonPagination) GetTotal() int  { return this.Total }
func (this SeasonPagination) GetFirst() int  { return this.First }
func (this SeasonPagination) GetOffset() int { return this.Offset }

type SeasonSearchItem struct {
	ID          string  `json:"id"`
	LegacyID    *string `json:"legacyID"`
	AgeRating   string  `json:"ageRating"`
	Collection  string  `json:"collection"`
	Title       string  `json:"title"`
	Header      *string `json:"header"`
	Description *string `json:"description"`
	Highlight   *string `json:"highlight"`
	Image       *string `json:"image"`
	URL         string  `json:"url"`
	ShowID      string  `json:"showId"`
	ShowTitle   string  `json:"showTitle"`
	Show        *Show   `json:"show"`
}

func (SeasonSearchItem) IsSearchResultItem()          {}
func (this SeasonSearchItem) GetID() string           { return this.ID }
func (this SeasonSearchItem) GetLegacyID() *string    { return this.LegacyID }
func (this SeasonSearchItem) GetCollection() string   { return this.Collection }
func (this SeasonSearchItem) GetTitle() string        { return this.Title }
func (this SeasonSearchItem) GetHeader() *string      { return this.Header }
func (this SeasonSearchItem) GetDescription() *string { return this.Description }
func (this SeasonSearchItem) GetHighlight() *string   { return this.Highlight }
func (this SeasonSearchItem) GetImage() *string       { return this.Image }
func (this SeasonSearchItem) GetURL() string          { return this.URL }

type SectionItem struct {
	ID          string          `json:"id"`
	Sort        int             `json:"sort"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Image       *string         `json:"image"`
	Item        SectionItemType `json:"item"`
}

type SectionItemPagination struct {
	First  int            `json:"first"`
	Offset int            `json:"offset"`
	Total  int            `json:"total"`
	Items  []*SectionItem `json:"items"`
}

func (SectionItemPagination) IsPagination()       {}
func (this SectionItemPagination) GetTotal() int  { return this.Total }
func (this SectionItemPagination) GetFirst() int  { return this.First }
func (this SectionItemPagination) GetOffset() int { return this.Offset }

type SectionPagination struct {
	Total  int       `json:"total"`
	First  int       `json:"first"`
	Offset int       `json:"offset"`
	Items  []Section `json:"items"`
}

func (SectionPagination) IsPagination()       {}
func (this SectionPagination) GetTotal() int  { return this.Total }
func (this SectionPagination) GetFirst() int  { return this.First }
func (this SectionPagination) GetOffset() int { return this.Offset }

type Settings struct {
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
}

type Show struct {
	ID             string            `json:"id"`
	LegacyID       *string           `json:"legacyID"`
	Type           ShowType          `json:"type"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	Image          *string           `json:"image"`
	ImageURL       *string           `json:"imageUrl"`
	Images         []*Image          `json:"images"`
	EpisodeCount   int               `json:"episodeCount"`
	SeasonCount    int               `json:"seasonCount"`
	Seasons        *SeasonPagination `json:"seasons"`
	DefaultEpisode *Episode          `json:"defaultEpisode"`
}

func (Show) IsSectionItemType() {}

type ShowCalendarEntry struct {
	ID          string `json:"id"`
	Event       *Event `json:"event"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Show        *Show  `json:"show"`
}

func (ShowCalendarEntry) IsCalendarEntry()            {}
func (this ShowCalendarEntry) GetID() string          { return this.ID }
func (this ShowCalendarEntry) GetEvent() *Event       { return this.Event }
func (this ShowCalendarEntry) GetTitle() string       { return this.Title }
func (this ShowCalendarEntry) GetDescription() string { return this.Description }
func (this ShowCalendarEntry) GetStart() string       { return this.Start }
func (this ShowCalendarEntry) GetEnd() string         { return this.End }

type ShowItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Show     *Show    `json:"show"`
}

func (ShowItem) IsCollectionItem()         {}
func (this ShowItem) GetID() string        { return this.ID }
func (this ShowItem) GetSort() int         { return this.Sort }
func (this ShowItem) GetTitle() string     { return this.Title }
func (this ShowItem) GetImageURL() *string { return this.ImageURL }
func (this ShowItem) GetImages() []*Image {
	if this.Images == nil {
		return nil
	}
	interfaceSlice := make([]*Image, 0, len(this.Images))
	for _, concrete := range this.Images {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type ShowSearchItem struct {
	ID          string  `json:"id"`
	LegacyID    *string `json:"legacyID"`
	Collection  string  `json:"collection"`
	Title       string  `json:"title"`
	Header      *string `json:"header"`
	Description *string `json:"description"`
	Highlight   *string `json:"highlight"`
	Image       *string `json:"image"`
	URL         string  `json:"url"`
}

func (ShowSearchItem) IsSearchResultItem()          {}
func (this ShowSearchItem) GetID() string           { return this.ID }
func (this ShowSearchItem) GetLegacyID() *string    { return this.LegacyID }
func (this ShowSearchItem) GetCollection() string   { return this.Collection }
func (this ShowSearchItem) GetTitle() string        { return this.Title }
func (this ShowSearchItem) GetHeader() *string      { return this.Header }
func (this ShowSearchItem) GetDescription() *string { return this.Description }
func (this ShowSearchItem) GetHighlight() *string   { return this.Highlight }
func (this ShowSearchItem) GetImage() *string       { return this.Image }
func (this ShowSearchItem) GetURL() string          { return this.URL }

type SimpleCalendarEntry struct {
	ID          string `json:"id"`
	Event       *Event `json:"event"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

func (SimpleCalendarEntry) IsCalendarEntry()            {}
func (this SimpleCalendarEntry) GetID() string          { return this.ID }
func (this SimpleCalendarEntry) GetEvent() *Event       { return this.Event }
func (this SimpleCalendarEntry) GetTitle() string       { return this.Title }
func (this SimpleCalendarEntry) GetDescription() string { return this.Description }
func (this SimpleCalendarEntry) GetStart() string       { return this.Start }
func (this SimpleCalendarEntry) GetEnd() string         { return this.End }

type Stream struct {
	ID                string     `json:"id"`
	URL               string     `json:"url"`
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
	Type              StreamType `json:"type"`
}

type User struct {
	ID        *string   `json:"id"`
	Anonymous bool      `json:"anonymous"`
	BccMember bool      `json:"bccMember"`
	Audience  *string   `json:"audience"`
	Email     *string   `json:"email"`
	Settings  *Settings `json:"settings"`
	Roles     []string  `json:"roles"`
}

type WebSection struct {
	ID             string               `json:"id"`
	Metadata       *ItemSectionMetadata `json:"metadata"`
	Title          *string              `json:"title"`
	URL            string               `json:"url"`
	WidthRatio     float64              `json:"widthRatio"`
	Authentication bool                 `json:"authentication"`
}

func (WebSection) IsSection()             {}
func (this WebSection) GetID() string     { return this.ID }
func (this WebSection) GetTitle() *string { return this.Title }

type EpisodeType string

const (
	EpisodeTypeEpisode    EpisodeType = "episode"
	EpisodeTypeStandalone EpisodeType = "standalone"
)

var AllEpisodeType = []EpisodeType{
	EpisodeTypeEpisode,
	EpisodeTypeStandalone,
}

func (e EpisodeType) IsValid() bool {
	switch e {
	case EpisodeTypeEpisode, EpisodeTypeStandalone:
		return true
	}
	return false
}

func (e EpisodeType) String() string {
	return string(e)
}

func (e *EpisodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EpisodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EpisodeType", str)
	}
	return nil
}

func (e EpisodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GridSectionSize string

const (
	GridSectionSizeHalf GridSectionSize = "half"
)

var AllGridSectionSize = []GridSectionSize{
	GridSectionSizeHalf,
}

func (e GridSectionSize) IsValid() bool {
	switch e {
	case GridSectionSizeHalf:
		return true
	}
	return false
}

func (e GridSectionSize) String() string {
	return string(e)
}

func (e *GridSectionSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GridSectionSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GridSectionSize", str)
	}
	return nil
}

func (e GridSectionSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ImageStyle string

const (
	ImageStylePoster   ImageStyle = "poster"
	ImageStyleFeatured ImageStyle = "featured"
	ImageStyleDefault  ImageStyle = "default"
)

var AllImageStyle = []ImageStyle{
	ImageStylePoster,
	ImageStyleFeatured,
	ImageStyleDefault,
}

func (e ImageStyle) IsValid() bool {
	switch e {
	case ImageStylePoster, ImageStyleFeatured, ImageStyleDefault:
		return true
	}
	return false
}

func (e ImageStyle) String() string {
	return string(e)
}

func (e *ImageStyle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImageStyle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImageStyle", str)
	}
	return nil
}

func (e ImageStyle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Language string

const (
	LanguageEn Language = "en"
	LanguageNo Language = "no"
	LanguageDe Language = "de"
)

var AllLanguage = []Language{
	LanguageEn,
	LanguageNo,
	LanguageDe,
}

func (e Language) IsValid() bool {
	switch e {
	case LanguageEn, LanguageNo, LanguageDe:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SectionSize string

const (
	SectionSizeSmall  SectionSize = "small"
	SectionSizeMedium SectionSize = "medium"
)

var AllSectionSize = []SectionSize{
	SectionSizeSmall,
	SectionSizeMedium,
}

func (e SectionSize) IsValid() bool {
	switch e {
	case SectionSizeSmall, SectionSizeMedium:
		return true
	}
	return false
}

func (e SectionSize) String() string {
	return string(e)
}

func (e *SectionSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SectionSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SectionSize", str)
	}
	return nil
}

func (e SectionSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShowType string

const (
	ShowTypeEvent  ShowType = "event"
	ShowTypeSeries ShowType = "series"
)

var AllShowType = []ShowType{
	ShowTypeEvent,
	ShowTypeSeries,
}

func (e ShowType) IsValid() bool {
	switch e {
	case ShowTypeEvent, ShowTypeSeries:
		return true
	}
	return false
}

func (e ShowType) String() string {
	return string(e)
}

func (e *ShowType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShowType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShowType", str)
	}
	return nil
}

func (e ShowType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StreamType string

const (
	StreamTypeHlsTs   StreamType = "hls_ts"
	StreamTypeHlsCmaf StreamType = "hls_cmaf"
	StreamTypeDash    StreamType = "dash"
)

var AllStreamType = []StreamType{
	StreamTypeHlsTs,
	StreamTypeHlsCmaf,
	StreamTypeDash,
}

func (e StreamType) IsValid() bool {
	switch e {
	case StreamTypeHlsTs, StreamTypeHlsCmaf, StreamTypeDash:
		return true
	}
	return false
}

func (e StreamType) String() string {
	return string(e)
}

func (e *StreamType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StreamType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StreamType", str)
	}
	return nil
}

func (e StreamType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
