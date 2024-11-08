package response

type References struct {
	Reviews []Review `json:"reviews"`
	Places  []string `json:"places"`
}
