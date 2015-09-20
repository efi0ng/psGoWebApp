package controllers

import (
	"net/http"
	"viewmodels"
	"text/template"
)

type productsController struct {
	template *template.Template
}

func (this *productsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetProducts()
	
	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}


