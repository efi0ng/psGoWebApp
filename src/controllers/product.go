package controllers

import (
	"models"
	"controllers/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
	"converters"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	product, err := models.GetProductById(id)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetProductPage(product.Name())
	vm.Product = converters.ConvertProductToViewModel(product)
	this.template.Execute(responseWriter, vm)
}
