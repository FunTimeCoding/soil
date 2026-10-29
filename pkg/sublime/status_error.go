package sublime

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func statusError(
	operation string,
	r *http.Response,
) error {
	reason := refusal.Reason(system.ReadAll(r.Body), r.Status)

	switch r.StatusCode {
	case http.StatusNotFound:
		return not_found.Format("%s: %s", operation, reason)
	case http.StatusBadRequest:
		return validation.New("%s: %s", operation, reason)
	case http.StatusConflict:
		return conflict.Format("%s: %s", operation, reason)
	default:
		return unexpected.Format("%s: %d: %s", operation, r.StatusCode, reason)
	}
}
