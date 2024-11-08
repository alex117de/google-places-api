package response

type GenerativeSummary struct {
	Overview                  LocalizedText `json:"overview"`
	OverviewFlagContentURI    string        `json:"overviewFlagContentUri"`
	Description               LocalizedText `json:"description"`
	DescriptionFlagContentURI string        `json:"descriptionFlagContentUri"`
	References                References    `json:"references"`
}
