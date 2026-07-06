package web

import (
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func ReadBytes(r *http.Response) []byte {
	return system.ReadAll(r.Body)
}
