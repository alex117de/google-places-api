package response

type AreaSummary struct {
	ContentBlocks  []ContentBlock `json:"contentBlocks"`
	FlagContentURI string         `json:"flagContentUri"`
}
