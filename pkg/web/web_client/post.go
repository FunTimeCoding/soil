package web_client

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/web_client/web_response"
	"time"
)

func (c *Client) Post(
	locator string,
	body any,
) (*web_response.Response, error) {
	encoded := notation.Encode(body, false)
	start := c.clock.Now()
	response, e := c.client.Post(
		locator,
		constant.Object,
		bytes.NewBuffer([]byte(encoded)),
	)
	r := web_response.New(response, time.Since(start).Milliseconds())
	r.BodyBytes = len(encoded)

	return r, e
}
