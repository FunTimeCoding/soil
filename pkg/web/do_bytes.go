package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
)

func DoBytes(
	c *http.Client,
	r *http.Request,
) ([]byte, error) {
	response, e := Do(c, r)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(response.Body)

	return io.ReadAll(response.Body)
}
