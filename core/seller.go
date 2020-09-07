package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Seller struct {
	Name      string
	ContactNo string
	Email     string
	Province  string
	Town      string
	Suburb    string
}

func (s Seller) Valid() error {
	return validation.Struct(s)
}
