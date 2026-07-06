package errors

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
)

func PanicStatus(r *http.Response) {
	if r.StatusCode != http.StatusOK {
		unexpected.Integer(r.StatusCode)
	}
}
