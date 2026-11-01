package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostRun(
	x context.Context,
	intent string,
) error {
	r, e := c.client.PostRunWithResponse(x, client.RunRequest{Intent: intent})

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusAccepted {
		return refusal.New("run", r.HTTPResponse, r.Body)
	}

	return nil
}
