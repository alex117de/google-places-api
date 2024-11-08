package response

type ParkingOptions struct {
	FreeParkingLot    bool `json:"freeParkingLot"`
	PaidParkingLot    bool `json:"paidParkingLot"`
	FreeStreetParking bool `json:"freeStreetParking"`
	PaidStreetParking bool `json:"paidStreetParking"`
	ValetParking      bool `json:"valetParking"`
	FreeGarageParking bool `json:"freeGarageParking"`
	PaidGarageParking bool `json:"paidGarageParking"`
}
