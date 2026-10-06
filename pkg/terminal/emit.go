package terminal

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"net/http"
)

func (t *Terminal) Emit(r *response.Response) {
	if r.Status >= http.StatusBadRequest {
		t.Exitln(r.Body)

		return
	}

	if r.Body == "" {
		return
	}

	t.Line(r.Body)
}
