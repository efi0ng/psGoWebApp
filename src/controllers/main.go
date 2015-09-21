package controllers

import (
	"bufio"
	"controllers/util"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	router := mux.NewRouter()

	// handle templates
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	catController := new(categoryController)
	catController.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", catController.get)

	pcController := new(productController)
	pcController.template = templates.Lookup("product.html")
	router.HandleFunc("/product/{id}", pcController.get)

	profileController := new(profileController)
	profileController.template = templates.Lookup("profile.html")
	router.HandleFunc("/profile", profileController.handle)
	
	http.Handle("/", router)

	// handle resources
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/video/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".mp4") {
		contentType = "video/mp4"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	defer f.Close()
	w.Header().Add("Content-Type", contentType)

	// add gzip compression if possible
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	br := bufio.NewReader(f)
	br.WriteTo(responseWriter)
}
