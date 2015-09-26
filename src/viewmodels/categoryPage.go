package viewmodels

import ()

type CategoryPage struct {
	Title    string
	Active   string
	Products []Product
}

func GetCategoryPage(title string) CategoryPage {
	var result CategoryPage

	result.Active = "shop"
	result.Title = "Lemonade Stand Society - " + title + " Shop"

	return result
}
