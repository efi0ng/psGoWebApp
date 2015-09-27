package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

type homeController struct {
	template *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetHome()
	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {	
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetLogin()
	this.loginTemplate.Execute(responseWriter, vm)
}
