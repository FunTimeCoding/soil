package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"io"
	"net/http"
)

func (c *Client) PostPackage(
	x context.Context,
	version string,
	repository string,
	architecture string,
	file string,
	body io.Reader,
) error {
	r, e := c.client.PostPackageWithBodyWithResponse(
		x,
		version,
		repository,
		architecture,
		file,
		constant.Octets,
		body,
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusCreated {
		return refusal.New("post package", r.HTTPResponse, r.Body)
	}

	return nil
}
