package converters

import (
	"models"
	"viewmodels"
)

func ConvertCategoryToViewModel(category models.Category, isOrientRight bool) viewmodels.Category {
	result := viewmodels.Category{
		ImageUrl:      category.ImageUrl(),
		Title:         category.Title(),
		Description:   category.Description(),
		Id:            category.Id(),
		IsOrientRight: isOrientRight,
	}

	return result
}
