package response

type Photo struct {
	Name               string              `json:"name"`
	WidthPx            int                 `json:"widthPx"`
	HeightPx           int                 `json:"heightPx"`
	AuthorAttributions []AuthorAttribution `json:"authorAttributions"`
	FlagContentURI     string              `json:"flagContentUri"`
	GoogleMapsURI      string              `json:"googleMapsUri"`
}
