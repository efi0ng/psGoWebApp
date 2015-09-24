package converters

import (
	"models"
	"viewmodels"
)

func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		ImageUrl:      category.ImageUrl(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
		Id:            category.Id(),
	}

	return result
}
