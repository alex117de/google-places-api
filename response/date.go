package response

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type SpecialDay struct {
	Date Date `json:"date"`
}
