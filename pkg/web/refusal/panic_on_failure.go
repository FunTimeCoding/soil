package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func PanicOnFailure(
	operation string,
	r *http.Response,
	body []byte,
) {
	if r.StatusCode >= http.StatusOK &&
		r.StatusCode < http.StatusMultipleChoices {
		return
	}

	errors.PanicOnError(New(operation, r, body))
}
