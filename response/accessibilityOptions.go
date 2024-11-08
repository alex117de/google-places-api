package response

type AccessibilityOptions struct {
	WheelchairAccessibleParking  bool `json:"wheelchairAccessibleParking"`
	WheelchairAccessibleEntrance bool `json:"wheelchairAccessibleEntrance"`
	WheelchairAccessibleRestroom bool `json:"wheelchairAccessibleRestroom"`
	WheelchairAccessibleSeating  bool `json:"wheelchairAccessibleSeating"`
}
