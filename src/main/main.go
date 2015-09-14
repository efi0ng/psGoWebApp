package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
	"viewmodels"
)

func main() {
	templates := populateTemplates()

	// handle templates
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")

			var context interface{} = nil

			switch requestedFile {
			case "home":
				context = viewmodels.GetHome()
			case "categories":
				context = viewmodels.GetCategories()
			}

			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

	// handle resources
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/video/", serveResource)

	http.ListenAndServe(":8000", nil)
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

	br := bufio.NewReader(f)
	br.WriteTo(w)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
