package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
)

func Status(
	status int,
	reason string,
) error {
	switch status {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return validation.New("%s", reason)
	case http.StatusForbidden:
		return forbidden.Format("%s", reason)
	case http.StatusNotFound:
		return not_found.Format("%s", reason)
	case http.StatusConflict:
		return conflict.Format("%s", reason)
	case http.StatusTooManyRequests,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return unavailable.Format("%s", reason)
	default:
		return nil
	}
}
