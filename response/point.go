package response

type Point struct {
	Date      Date `json:"date"`
	Truncated bool `json:"truncated"`
	Day       int  `json:"day"`
	Hour      int  `json:"hour"`
	Minute    int  `json:"minute"`
}
