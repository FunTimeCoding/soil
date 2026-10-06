package client

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"net/http"
)

func (c *Client) fail(
	w http.ResponseWriter,
	e error,
	message string,
) {
	if c.reporter != nil && classify.Reportable(e) {
		c.reporter.CaptureException(e)
	}

	http.Error(w, classify.Message(e, message), http.StatusBadGateway)
}
