package viewmodels

import (
	"errors"
)

var (
	productList []Product
)

type Products struct {
	Title    string
	Active   string
	Products []Product
}

func GetProducts(id int) Products {
	var result Products
	result.Active = "shop"
	var shopName string
	switch id {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}
	result.Title = "Lemonade Stand Society - " + shopName + " Shop"

	if id == 1 {
		result.Products = GetProductList()
	}

	return result
}

type ProductVM struct {
	Title   string
	Active  string
	Product Product
}

func GetProduct(id int) (ProductVM, error) {
	var result ProductVM
	var err error

	result.Active = "shop"
	result.Title = "Lemonade Stand Society - Kiwi Juice"

	for _, p := range GetProductList() {
		if p.Id == id {
			result.Product = p
			break
		}
	}

	if result.Product.Id == 0 {
		err = errors.New("No such product")
	}
	
	return result, err
}

func GetProductList() []Product {
	if productList == nil {
		lemonJuice := MakeLemonJuiceProduct()
		appleJuice := MakeAppleJuiceProduct()
		watermelonJuice := MakeWatermelonJuiceProduct()
		kiwiJuice := MakeKiwiJuiceProduct()
		mangosteenJuice := MakeMangosteenJuiceProduct()
		orangeJuice := MakeOrangeJuiceProduct()
		pineappleJuice := MakePineappleJuiceProduct()
		strawberryJuice := MakeStrawberryJuiceProduct()

		productList = []Product{
			lemonJuice,
			appleJuice,
			watermelonJuice,
			kiwiJuice,
			mangosteenJuice,
			orangeJuice,
			pineappleJuice,
			strawberryJuice,
		}
	}
	return productList
}
