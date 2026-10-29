package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func classifyRefusal(
	status int,
	body []byte,
) error {
	reason := refusal.Reason(body, http.StatusText(status))

	switch status {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return validation.New("linkace: %s", reason)
	case http.StatusUnauthorized, http.StatusForbidden:
		return forbidden.Format("linkace: %s", reason)
	case http.StatusConflict:
		return conflict.Format("linkace: %s", reason)
	case http.StatusServiceUnavailable:
		return unavailable.Format("linkace: %s", reason)
	default:
		return nil
	}
}
