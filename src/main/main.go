package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var path string

	if req.URL.Path == "/" {
		path = "public/html/home.html"
	} else {
		path = "public/" + req.URL.Path
	}

	data, err := ioutil.ReadFile(string(path))

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
		return
	}

	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".mp4") {
		contentType = "video/mpeg"
	} else {
		contentType = "text/plain"
	}
	
	w.Header().Add("Content-Type", contentType)
	w.Write(data)
}
