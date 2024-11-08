package response

type OpeningHours struct {
	Periods             []Period     `json:"periods"`
	WeekdayDescriptions []string     `json:"weekdayDescriptions"`
	SecondaryHoursType  string       `json:"secondaryHoursType"` // could be an enum if defined
	SpecialDays         []SpecialDay `json:"specialDays"`
	NextOpenTime        string       `json:"nextOpenTime"`
	NextCloseTime       string       `json:"nextCloseTime"`
	OpenNow             bool         `json:"openNow"`
}
