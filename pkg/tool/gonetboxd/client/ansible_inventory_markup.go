package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) AnsibleInventoryMarkup(
	x context.Context,
	field string,
	site string,
) (string, error) {
	p := &client.GetAnsibleInventoryParams{Field: &field}

	if site != "" {
		p.Site = &site
	}

	r, e := c.client.GetAnsibleInventoryWithResponse(x, p)

	if e != nil {
		return "", e
	}

	if r.StatusCode() != http.StatusOK {
		return "", refusal.New("ansible inventory", r.HTTPResponse, r.Body)
	}

	return string(r.Body), nil
}
