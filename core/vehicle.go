package core

import "github.com/louisevanderlith/husk"

type Vehicle struct {
	Year           int
	Make           string
	MakeCountry    string
	Model          string
	Trim           string
	Drive          string
	Transmission   string
	Body           string
	EnginePosition string
	Mileage        int
	Price          float32
	Condition      int
	Issues         string
}

func (v Vehicle) Valid() error {
	return husk.ValidateStruct(&v)
}
