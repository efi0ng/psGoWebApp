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
		container := template.New("container")

		_, err := container.New("main").Parse(doc)
		if err != nil {
			return
		}

		_, err = container.New("header").Parse(header)
		if err != nil {
			return
		}

		_, err = container.New("footer").Parse(footer)
		if err != nil {
			return
		}

		context := Context{
			"Fruit Store Galore",
			"Jake Jones",
			req.URL.Path,
			[4]string{"Lemon", "Orange", "Apple", "Pear"},
		}
		container.Lookup("main").Execute(w, context)
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `
{{template "header" .Title}}
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
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
	<head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type Context struct {
	Title   string
	Message string
	Path    string
	Fruit   [4]string
}

func (this Context) Knowledge(tidbit string) string {
	return fmt.Sprint("Wise wisdom (", tidbit, ") at ", time.Now().Format(time.UnixDate))
}
