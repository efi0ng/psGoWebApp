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
		result.Products = getProductList()
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

	var product Product
	for _, p := range getProductList() {
		if p.Id == id {
			product = p
			break
		}
	}

	if product.Id == 0 {
		err = errors.New("No such product")
		return result, err
	}
	
	result.Active = "shop"
	result.Product = product
	result.Title = "Lemonade Stand Society - " + result.Product.Name
	
	return result, err
}

func getProductList() []Product {
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
