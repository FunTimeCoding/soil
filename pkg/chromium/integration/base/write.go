package base

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
)

func write(
	w http.ResponseWriter,
	body string,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Print(w, "%s", body)
}
