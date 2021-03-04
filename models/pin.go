package models

import "github.com/BrandonRomano/iso8601"

const PIN_FIELDS = "id,link,url,board,created_at,note,color,counts,media,attribution,image,metadata"

type Pin struct {
	Id          string       `json:"id"`
	Link        string       `json:"link"`
	Url         string       `json:"url"`
	Board       Board        `json:"board"`
	CreatedAt   iso8601.Time `json:"created_at"`
	Note        string       `json:"note"`
	Color       string       `json:"color"`
	Counts      PinCounts    `json:"counts"`
	Media       Media        `json:"media"`
	Attribution Attribution  `json:"attribution"`
	Image       PinImage     `json:"image"`
	Metadata    PinMetadata  `json:"metadata"`
}

type PinImage struct {
	Original Image `json:"original"`
}

type PinCounts struct {
	Likes    int32 `json:"likes"`
	Comments int32 `json:"comments"`
	Repins   int32 `json:"repins"`
}

type Media struct {
	Type string `json:"type"`
}

type Attribution struct {
	Title              string `json:"title"`
	Url                string `json:"url"`
	ProviderIconUrl    string `json:"provider_icon_url"`
	AuthorName         string `json:"author_name"`
	ProviderFaviconUrl string `json:"provider_favicon_url"`
	AuthorUrl          string `json:"author_url"`
	ProviderName       string `json:"provider_name"`
}
