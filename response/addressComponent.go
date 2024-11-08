package response

type AddressComponent struct {
	LongText     string   `json:"longText"`
	ShortText    string   `json:"shortText"`
	Types        []string `json:"types"`
	LanguageCode string   `json:"languageCode"`
}
