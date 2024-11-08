package response

type FuelPrice struct {
	Type       string `json:"type"` // could be an enum if defined
	Price      Money  `json:"price"`
	UpdateTime string `json:"updateTime"`
}
