package refusing

import (
	"context"
	"net/http"
)

func (c *Client) Do(
	x context.Context,
	r *http.Request,
) (*http.Response, []byte, error) {
	response, body, e := c.client.Do(x, r)

	if e != nil {
		return response, body, e
	}

	return response, body, classify(response.StatusCode, body)
}
