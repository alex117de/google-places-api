package response

type PriceRange struct {
	StartPrice Money `json:"startPrice"`
	EndPrice   Money `json:"endPrice"`
}
