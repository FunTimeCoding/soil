package console

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func Emit(r *response.Response) {
	if r.Status >= http.StatusBadRequest {
		system.Exitln(r.Body)
	}

	if r.Body == "" {
		return
	}

	Line(r.Body)
}
