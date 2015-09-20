package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"
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

	vm, err := viewmodels.GetProduct(id)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	
	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
