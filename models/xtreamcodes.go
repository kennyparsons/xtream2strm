package models

// VODStream represents the information of a single VOD stream.
type VODStream struct {
	Added              string      `json:"added"`
	CategoryID         string      `json:"category_id"`
	ContainerExtension string      `json:"container_extension"`
	CustomSID          string      `json:"custom_sid"`
	DirectSource       string      `json:"direct_source"`
	Name               string      `json:"name"`
	Num                int         `json:"num"`
	Rating             interface{} `json:"rating"`
	Rating5Based       float64     `json:"rating_5based"`
	StreamIcon         string      `json:"stream_icon"`
	StreamID           int         `json:"stream_id"`
	StreamType         string      `json:"stream_type"`
}

// XtreamCodesJSON represents the JSON response from the Xtream Codes API for VOD streams.
type XtreamCodesJSON []VODStream

// Series represents the information of a single series.
type Series struct {
	BackdropPath   []string `json:"backdrop_path"`
	Cast           string   `json:"cast"`
	CategoryID     string   `json:"category_id"`
	Cover          string   `json:"cover"`
	Director       string   `json:"director"`
	EpisodeRunTime string   `json:"episode_run_time"`
	Genre          string   `json:"genre"`
	LastModified   string   `json:"last_modified"`
	Name           string   `json:"name"`
	Num            int      `json:"num"`
	Plot           string   `json:"plot"`
	Rating         string   `json:"rating"`
	Rating5Based   float64  `json:"rating_5based"`
	ReleaseDate    string   `json:"releaseDate"`
	SeriesID       int      `json:"series_id"`
	YoutubeTrailer string   `json:"youtube_trailer"`
}

// SeriesJSON represents the JSON response from the Xtream Codes API for series.
type SeriesJSON []Series

type EpisodeInfo struct {
	ID                 string `json:"id"`
	EpisodeNum         int    `json:"episode_num"`
	Season             int    `json:"season"`
	Title              string `json:"title"`
	ContainerExtension string `json:"container_extension"`
	// Add other fields as needed
}

type Season struct {
	SeasonNumber int `json:"season_number"`
	// Add other fields as needed
}

type SeriesInfoResponse struct {
	Episodes map[string][]EpisodeInfo `json:"episodes"`
	Seasons  []Season                 `json:"seasons"`
	// Add other fields as needed
}
