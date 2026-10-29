package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
)

func New(
	operation string,
	r *http.Response,
	body []byte,
) error {
	reason := Reason(body, r.Status)

	switch r.StatusCode {
	case http.StatusBadRequest:
		return validation.New("%s: %s", operation, reason)
	case http.StatusForbidden:
		return forbidden.Format("%s: %s", operation, reason)
	case http.StatusNotFound:
		return not_found.Format("%s: %s", operation, reason)
	case http.StatusConflict:
		return conflict.Format("%s: %s", operation, reason)
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
