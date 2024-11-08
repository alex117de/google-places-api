package response

type LocalizedText struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}
