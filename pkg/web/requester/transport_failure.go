package requester

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"net/http"
	"net/url"
)

func transportFailure(
	h *http.Request,
	e error,
) error {
	if f := connection.Classify(e); f != nil {
		return f
	}

	if transport, failed := errors.AsType[*url.Error](e); failed {
		e = transport.Err
	}

	return fmt.Errorf("send %s%s: %w", h.URL.Host, h.URL.Path, e)
}
