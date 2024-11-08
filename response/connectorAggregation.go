package response

type ConnectorAggregation struct {
	Type                       string  `json:"type"` // could be an enum if defined
	MaxChargeRateKw            float64 `json:"maxChargeRateKw"`
	Count                      int     `json:"count"`
	AvailabilityLastUpdateTime string  `json:"availabilityLastUpdateTime"`
	AvailableCount             int     `json:"availableCount"`
	OutOfServiceCount          int     `json:"outOfServiceCount"`
}
