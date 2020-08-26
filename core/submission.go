package core

import "github.com/louisevanderlith/husk"

type Submission struct {
	Seller     Seller
	VehicleKey husk.Key
	Price	float32
}

func (s Submission) Valid() error {
	return husk.ValidateStruct(s)
}
