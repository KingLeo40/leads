package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Submission struct {
	Seller     Seller
	VehicleKey hsk.Key
	Price      float32
}

func (s Submission) Valid() error {
	return validation.Struct(s)
}
