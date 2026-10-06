package web

import (
	"errors"
	"fmt"
	libraryErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
	"net/url"
)

func Do(
	c *http.Client,
	r *http.Request,
) (*http.Response, error) {
	result, e := c.Do(r)

	if f := connection.Classify(e); f != nil {
		return nil, f
	}

	if transport, failed := errors.AsType[*url.Error](e); failed {
		return nil, fmt.Errorf(
			"send %s%s: %w",
			r.URL.Host,
			r.URL.Path,
			transport.Err,
		)
	}

	if e != nil {
		return nil, e
	}

	if ResponseOkay(result) {
		return result, nil
	}

	libraryErrors.LogClose(result.Body)

	if result.StatusCode == http.StatusNotFound {
		return nil, not_found.Format(
			"resource not found: %s%s",
			r.URL.Host,
			r.URL.Path,
		)
	}

	return nil, unexpected.Format(
		"%s%s status: %d",
		r.URL.Host,
		r.URL.Path,
		result.StatusCode,
	)
}
