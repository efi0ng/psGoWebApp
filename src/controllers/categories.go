package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	vm := viewmodels.GetProducts(id)

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
