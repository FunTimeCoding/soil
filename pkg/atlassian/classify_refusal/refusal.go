package classify_refusal

import (
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func Refusal(
	status int,
	body []byte,
) error {
	if status == http.StatusBadRequest {
		return nil
	}

	return refusal.Status(status, refusal.Reason(body, http.StatusText(status)))
}
