package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetProfile()
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}
