package controllers

import (
	"net/http"
	"viewmodels"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()
	
	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}

