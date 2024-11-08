package response

type GoogleMapsLinks struct {
	DirectionsURI   string `json:"directionsUri"`
	PlaceURI        string `json:"placeUri"`
	WriteAReviewURI string `json:"writeAReviewUri"`
	ReviewsURI      string `json:"reviewsUri"`
	PhotosURI       string `json:"photosUri"`
}
