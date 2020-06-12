package core

import "github.com/louisevanderlith/husk"

type Submission struct {
	Seller Seller
	VehicleKey husk.Key
	Photos Photos
}

func (s Submission) Valid() (bool, error){
	return husk.ValidateStruct(&s)
}