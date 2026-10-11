package client_tester

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
)

func user(
	w http.ResponseWriter,
	_ *http.Request,
) {
	writer.Print(w, `{"id": 1, "username": "tester"}`)
}
