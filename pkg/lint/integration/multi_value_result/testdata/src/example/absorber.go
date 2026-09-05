package example

import "text/template"

var Home = template.Must(template.New("home").Parse("<p>{{.}}</p>"))
