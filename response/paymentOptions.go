package response

type PaymentOptions struct {
	AcceptsCreditCards bool `json:"acceptsCreditCards"`
	AcceptsDebitCards  bool `json:"acceptsDebitCards"`
	AcceptsCashOnly    bool `json:"acceptsCashOnly"`
	AcceptsNFC         bool `json:"acceptsNfc"`
}
