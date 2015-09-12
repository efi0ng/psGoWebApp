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

		context := Context{
			"Fruit Store Galore",
			"Jake Jones",
			req.URL.Path,
			[3]string{"Lemon", "Orange", "Apple"},
		}
		tmpl.Execute(w, context)
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
	<head><title>{{.Title}}</title></head>
	<body>
		{{if eq .Path "/Google"}}
			<h1>Hey, Google made Go!</h1>
		{{else}} 
			<h1>Hola {{.Message}}!</h1>
		{{end}}
		<p>You have reached {{.Path}}. {{.Knowledge "hope"}}</p>
		<ul>
		{{range .Fruit}}
		  <li>{{.}}</li>
		{{else}}
			<li>No fruit found?</li>
		{{end}}
		</ul>
	</body>
</html>
`

type Context struct {
	Title   string
	Message string
	Path    string
	Fruit   [3]string
}

func (this Context) Knowledge(tidbit string) string {
	return fmt.Sprint("Wise wisdom (", tidbit, ") at ", time.Now().Format(time.UnixDate))
}
