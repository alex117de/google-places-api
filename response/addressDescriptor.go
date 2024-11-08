package response

type AddressDescriptor struct {
	Landmarks []Landmark `json:"landmarks"`
	Areas     []Area     `json:"areas"`
}
