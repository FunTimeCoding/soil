package refusing

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func classify(
	status int,
	body []byte,
) error {
	if status/100 == 2 {
		return nil
	}

	reason := refusal.Reason(body, http.StatusText(status))

	switch status {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return validation.New("prometheus: %s", reason)
	default:
		return unexpected.Format("prometheus status: %d: %s", status, reason)
	}
}
