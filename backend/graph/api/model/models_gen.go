// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CalendarEntry interface {
	IsCalendarEntry()
}

type IconItem interface {
	IsIconItem()
}

type Item interface {
	IsItem()
}

type ItemSection interface {
	Section
	IsItemSection()
}

type LabelItem interface {
	IsLabelItem()
}

type LinkSection interface {
	Section
	IsLinkSection()
}

type Pagination interface {
	IsPagination()
}

type SearchResultItem interface {
	IsSearchResultItem()
}

type Section interface {
	IsSection()
}

type SectionItemType interface {
	IsSectionItemType()
}

type Application struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	ClientVersion string `json:"clientVersion"`
	Page          *Page  `json:"page"`
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
	Total  int    `json:"total"`
	First  int    `json:"first"`
	Offset int    `json:"offset"`
	Items  []Item `json:"items"`
}

func (CollectionItemPagination) IsPagination() {}

type Config struct {
	Global *GlobalConfig `json:"global"`
}

type DefaultSection struct {
	ID    string                 `json:"id"`
	Title *string                `json:"title"`
	Size  SectionSize            `json:"size"`
	Items *SectionItemPagination `json:"items"`
}

func (DefaultSection) IsSection()     {}
func (DefaultSection) IsItemSection() {}

type Episode struct {
	ID                string     `json:"id"`
	LegacyID          *string    `json:"legacyID"`
	LegacyProgramID   *string    `json:"legacyProgramID"`
	AgeRating         string     `json:"ageRating"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	ExtraDescription  string     `json:"extraDescription"`
	ImageURL          *string    `json:"imageUrl"`
	Streams           []*Stream  `json:"streams"`
	Files             []*File    `json:"files"`
	Chapters          []*Chapter `json:"chapters"`
	Season            *Season    `json:"season"`
	Duration          int        `json:"duration"`
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
	Images            []*Image   `json:"images"`
	Number            *int       `json:"number"`
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

func (EpisodeCalendarEntry) IsCalendarEntry() {}

type EpisodeItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Episode  *Episode `json:"episode"`
}

func (EpisodeItem) IsItem() {}

type EpisodePagination struct {
	Total  int        `json:"total"`
	First  int        `json:"first"`
	Offset int        `json:"offset"`
	Items  []*Episode `json:"items"`
}

func (EpisodePagination) IsPagination() {}

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

func (EpisodeSearchItem) IsSearchResultItem() {}

type Event struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
	Image string `json:"image"`
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

func (FAQCategoryPagination) IsPagination() {}

type FeaturedSection struct {
	ID    string                 `json:"id"`
	Title *string                `json:"title"`
	Size  SectionSize            `json:"size"`
	Items *SectionItemPagination `json:"items"`
}

func (FeaturedSection) IsSection()     {}
func (FeaturedSection) IsItemSection() {}

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

type GridSection struct {
	ID    string                 `json:"id"`
	Title *string                `json:"title"`
	Size  GridSectionSize        `json:"size"`
	Items *SectionItemPagination `json:"items"`
}

func (GridSection) IsSection()     {}
func (GridSection) IsItemSection() {}

type IconSection struct {
	ID    string  `json:"id"`
	Title *string `json:"title"`
}

func (IconSection) IsSection()     {}
func (IconSection) IsLinkSection() {}

type Image struct {
	Style string `json:"style"`
	URL   string `json:"url"`
}

type LabelItemPagination struct {
	Total  int         `json:"total"`
	First  int         `json:"first"`
	Offset int         `json:"offset"`
	Items  []LabelItem `json:"items"`
}

func (LabelItemPagination) IsPagination() {}

type LabelSection struct {
	ID    string               `json:"id"`
	Title *string              `json:"title"`
	Items *LabelItemPagination `json:"items"`
}

func (LabelSection) IsSection()     {}
func (LabelSection) IsLinkSection() {}

type MaintenanceMessage struct {
	Message string  `json:"message"`
	Details *string `json:"details"`
}

type Messages struct {
	Maintenance []*MaintenanceMessage `json:"maintenance"`
}

type Page struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       string             `json:"title"`
	Description *string            `json:"description"`
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

func (PageItem) IsItem() {}

type PageLabelItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Page  *Page  `json:"page"`
}

func (PageLabelItem) IsLabelItem() {}

type PosterSection struct {
	ID    string                 `json:"id"`
	Title *string                `json:"title"`
	Size  SectionSize            `json:"size"`
	Items *SectionItemPagination `json:"items"`
}

func (PosterSection) IsSection()     {}
func (PosterSection) IsItemSection() {}

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

func (QuestionPagination) IsPagination() {}

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

func (SeasonCalendarEntry) IsCalendarEntry() {}

type SeasonItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Season   *Season  `json:"season"`
}

func (SeasonItem) IsItem() {}

type SeasonPagination struct {
	Total  int       `json:"total"`
	First  int       `json:"first"`
	Offset int       `json:"offset"`
	Items  []*Season `json:"items"`
}

func (SeasonPagination) IsPagination() {}

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

func (SeasonSearchItem) IsSearchResultItem() {}

type SectionItem struct {
	ID    string          `json:"id"`
	Sort  int             `json:"sort"`
	Title string          `json:"title"`
	Image *string         `json:"image"`
	Item  SectionItemType `json:"item"`
}

type SectionItemPagination struct {
	First  int            `json:"first"`
	Offset int            `json:"offset"`
	Total  int            `json:"total"`
	Items  []*SectionItem `json:"items"`
}

func (SectionItemPagination) IsPagination() {}

type SectionPagination struct {
	Total  int       `json:"total"`
	First  int       `json:"first"`
	Offset int       `json:"offset"`
	Items  []Section `json:"items"`
}

func (SectionPagination) IsPagination() {}

type Settings struct {
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
}

type Show struct {
	ID           string            `json:"id"`
	LegacyID     *string           `json:"legacyID"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	ImageURL     *string           `json:"imageUrl"`
	Images       []*Image          `json:"images"`
	EpisodeCount int               `json:"episodeCount"`
	SeasonCount  int               `json:"seasonCount"`
	Seasons      *SeasonPagination `json:"seasons"`
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

func (ShowCalendarEntry) IsCalendarEntry() {}

type ShowItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	Show     *Show    `json:"show"`
}

func (ShowItem) IsItem() {}

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

func (ShowSearchItem) IsSearchResultItem() {}

type SimpleCalendarEntry struct {
	ID          string `json:"id"`
	Event       *Event `json:"event"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

func (SimpleCalendarEntry) IsCalendarEntry() {}

type Stream struct {
	ID                string     `json:"id"`
	URL               string     `json:"url"`
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
	Type              StreamType `json:"type"`
}

type URLItem struct {
	ID       string   `json:"id"`
	Sort     int      `json:"sort"`
	Title    string   `json:"title"`
	ImageURL *string  `json:"imageUrl"`
	Images   []*Image `json:"images"`
	URL      string   `json:"url"`
}

func (URLItem) IsItem() {}

type URLLabelItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (URLLabelItem) IsLabelItem() {}

type User struct {
	ID        *string   `json:"id"`
	Anonymous bool      `json:"anonymous"`
	BccMember bool      `json:"bccMember"`
	Audience  *string   `json:"audience"`
	Email     *string   `json:"email"`
	Settings  *Settings `json:"settings"`
	Roles     []string  `json:"roles"`
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