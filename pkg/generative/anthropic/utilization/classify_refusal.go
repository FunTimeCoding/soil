package utilization

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func classifyRefusal(
	status int,
	body []byte,
) error {
	switch status {
	case http.StatusTooManyRequests,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return unavailable.Format(
			"utilization: %s",
			refusal.Reason(body, http.StatusText(status)),
		)
	default:
		return nil
	}
}
