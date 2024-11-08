package response

type EVChargeOptions struct {
	ConnectorCount       int                    `json:"connectorCount"`
	ConnectorAggregation []ConnectorAggregation `json:"connectorAggregation"`
}
