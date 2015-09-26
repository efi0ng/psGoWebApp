package viewmodels

import ()

type ProductPage struct {
	Title   string
	Active  string
	Product Product
}

func GetProductPage(name string) ProductPage {
	var result ProductPage

	result.Active = "shop"
	result.Title = "Lemonade Stand Society - " + name

	return result
}
