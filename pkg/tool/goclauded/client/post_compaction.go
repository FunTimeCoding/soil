package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostCompaction(
	x context.Context,
	body client.PostCompactionJSONRequestBody,
) error {
	r, e := c.generated.PostCompactionWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("post compaction", r.HTTPResponse, r.Body)
	}

	return nil
}
