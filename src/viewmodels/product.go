package viewmodels

import ()

type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLitre    float32
	PricePer10Litre  float32
	Origin           string
	IsOrganic        bool
	ImageUrl         string
	Id               int
}
