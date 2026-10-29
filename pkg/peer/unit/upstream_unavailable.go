package unit

import (
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func upstreamUnavailable() error {
	return refusal.New(
		"list issues",
		&http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Status:     "503 Service Unavailable",
		},
		nil,
	)
}
