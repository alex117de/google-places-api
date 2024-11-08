package response

type Area struct {
	Name        string        `json:"name"`
	PlaceID     string        `json:"placeId"`
	DisplayName LocalizedText `json:"displayName"`
	Containment string        `json:"containment"` // could be an enum if defined
}
