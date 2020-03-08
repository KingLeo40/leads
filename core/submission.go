package core

import "github.com/louisevanderlith/husk"

type Submission struct {
	Seller Seller
	Vehicle Vehicle
}

func (s Submission) Valid() (bool, error){
	return husk.ValidateStruct(&s)
}