package response

type Viewport struct {
	Low  LatLng `json:"low"`
	High LatLng `json:"high"`
}
