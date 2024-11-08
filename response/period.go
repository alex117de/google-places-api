package response

type Period struct {
	Open  Point `json:"open"`
	Close Point `json:"close"`
}
