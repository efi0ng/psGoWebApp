package controllers

import (
	"controllers/util"
	"models"
	"net/http"
	"text/template"
	"viewmodels"
)

type loginController struct {
	template *template.Template
}

func (this *loginController) login(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		_, _ = models.GetMember(email, password)
	}
	vm := viewmodels.GetLogin()
	this.template.Execute(responseWriter, vm)
}
