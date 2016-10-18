package reddit

type Media struct {
	Oembed Oembed `json:"oembed"`
	Type   string `json:"type"`
}

type Oembed struct {
	Description     string `json:"description"`
	Html            string `json:"html"`
	Height          string `json:"height"`
	ProviderName    string `json:"provider_name"`
	ProviderUrl     string `json:"provider_url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	ThumbnailUrl    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	Version         string `json:"version"`
	Width           int    `json:"width"`
}
