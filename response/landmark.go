package response

type Landmark struct {
	Name                       string        `json:"name"`
	PlaceID                    string        `json:"placeId"`
	DisplayName                LocalizedText `json:"displayName"`
	Types                      []string      `json:"types"`
	SpatialRelationship        string        `json:"spatialRelationship"` // could be an enum if defined
	StraightLineDistanceMeters float64       `json:"straightLineDistanceMeters"`
	TravelDistanceMeters       float64       `json:"travelDistanceMeters"`
}
