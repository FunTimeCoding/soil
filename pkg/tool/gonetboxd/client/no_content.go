package client

import (
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func noContent(
	operation string,
	r *http.Response,
	body []byte,
) error {
	if r.StatusCode != http.StatusNoContent {
		return refusal.New(operation, r, body)
	}

	return nil
}
