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
