package converters

import (
	"testing"
	"models"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetImageUrl("the image URL")
	category.SetTitle("the title")
	category.SetDescription("the description")
	category.SetId(42)
	isOrientRight := true
	
	result := ConvertCategoryToViewModel(category, isOrientRight)
	
	if result.ImageUrl != category.ImageUrl() {
		t.Log("Image URL not converted properly")
		t.Fail()
	}

	if result.Title != category.Title() {
		t.Log("Title not converted properly")
		t.Fail()
	}

	if result.Description != category.Description() {
		t.Log("Description not converted properly")
		t.Fail()
	}

	if result.IsOrientRight != isOrientRight {
		t.Log("IsOrientRight not converted properly")
		t.Fail()
	}

	if result.Id != category.Id() {
		t.Log("Id not converted properly")
		t.Fail()
	}

}
