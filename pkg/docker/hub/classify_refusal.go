package hub

import (
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func classifyRefusal(
	status int,
	body []byte,
) error {
	reason := refusal.Reason(body, http.StatusText(status))

	switch status {
	case http.StatusForbidden:
		return forbidden.Format("docker hub: %s", reason)
	case http.StatusTooManyRequests,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return unavailable.Format("docker hub: %s", reason)
	default:
		return nil
	}
}
