package iterm

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
)

func statusError(
	operation string,
	r *http.Response,
) error {
	switch r.StatusCode {
	case http.StatusNotFound:
		return not_found.Format("%s: %s", operation, reason(r))
	case http.StatusUnprocessableEntity:
		return validation.New("%s: %s", operation, reason(r))
	case http.StatusServiceUnavailable:
		return unavailable.Format("%s: %s", operation, reason(r))
	default:
		return unexpected.Format(
			"%s: %d: %s",
			operation,
			r.StatusCode,
			reason(r),
		)
	}
}
