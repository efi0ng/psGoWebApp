package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err != nil {
			return
		}

		context := Context{"Reader", req.URL.Path}
		tmpl.Execute(w, context)
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		<h1>Hello {{.Message}}</h1>
		<p>You tried to access {{.Path}}</p>
		<p>{{.Knowledge}}</p>
	</body>
</html>
`

type Context struct {
	Message string
	Path    string
}

func (this Context) Knowledge() string {
	return fmt.Sprint("Wise wisdom", time.Now)
}
