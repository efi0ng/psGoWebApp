package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"encoding/json"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(vm)
	if (err != nil) {
		responseWriter.WriteHeader(404)
		return
	}
	
	responseWriter.Write(data)
}
