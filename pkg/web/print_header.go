package web

import (
	"github.com/funtimecoding/soil/pkg/console"
	"net/http"
)

func PrintHeader(h http.Header) {
	for k, v := range h {
		console.Format("Header: %s: %s\n", k, v)
	}
}
