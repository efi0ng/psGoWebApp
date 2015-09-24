package controllers

import (
	"controllers/util"
	"converters"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()

	categoriesVM := []viewmodels.Category{}
	isOrientRight := false
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoryToViewModel(category, isOrientRight))
		isOrientRight = !isOrientRight
	}
	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
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
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}
