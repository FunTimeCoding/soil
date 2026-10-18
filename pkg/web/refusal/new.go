package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
)

func New(
	operation string,
	r *http.Response,
	body []byte,
) error {
	reason := Reason(body, r.Status)

	switch r.StatusCode {
	case http.StatusNotFound:
		return not_found.Format("%s: %s", operation, reason)
	case http.StatusServiceUnavailable:
		return unavailable.Format("%s: %s", operation, reason)
	}

	return unexpected.Format(
		"%s status: %d: %s",
		operation,
		r.StatusCode,
		reason,
	)
}
