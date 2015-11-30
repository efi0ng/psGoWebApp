package converters

import (
	"models"
	"viewmodels"
)

func ConvertProductToViewModel(product models.Product) viewmodels.Product {
	result := viewmodels.Product{
		Name: product.Name(),
		DescriptionShort: product.DescriptionShort(),
		DescriptionLong: product.DescriptionLong(),
		PricePerLitre: product.PricePerLitre(),
		PricePer10Litre: product.PricePer10Litre(),
		Origin: product.Origin(),
		IsOrganic: product.IsOrganic(),
		ImageUrl: product.ImageUrl(),
		Id: product.Id(),
	}
	
	return result
}