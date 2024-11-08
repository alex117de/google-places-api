package response

type Review struct {
	Name                           string            `json:"name"`
	RelativePublishTimeDescription string            `json:"relativePublishTimeDescription"`
	Text                           LocalizedText     `json:"text"`
	OriginalText                   LocalizedText     `json:"originalText"`
	Rating                         float64           `json:"rating"`
	AuthorAttribution              AuthorAttribution `json:"authorAttribution"`
	PublishTime                    string            `json:"publishTime"`
	FlagContentURI                 string            `json:"flagContentUri"`
	GoogleMapsURI                  string            `json:"googleMapsUri"`
}
