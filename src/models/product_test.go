package models

import (
	"testing"
)

func Test_GetJuiceProducts_ReturnsNonEmptySlice(t *testing.T) {
	products := getJuiceProducts()
	
	if len(products) == 0 {
		t.Log("GetJuiceProducts returned empty slice")
		t.Fail()
	}
}