package netbox

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"net/http"
)

func (c *Client) rawRequest(
	method string,
	path string,
	body any,
	result any,
) error {
	configuration := c.client.GetConfig()
	var reader io.Reader

	if body != nil {
		encoded, e := json.Marshal(body)

		if e != nil {
			return e
		}

		reader = bytes.NewReader(encoded)
	}

	q, e := http.NewRequestWithContext(
		c.context,
		method,
		join.Empty(configuration.Servers[0].URL, path),
		reader,
	)

	if e != nil {
		return e
	}

	for k, v := range configuration.DefaultHeader {
		q.Header.Set(k, v)
	}

	q.Header.Set("Content-Type", "application/json")
	r, e := configuration.HTTPClient.Do(q)

	if e != nil {
		return e
	}

	defer errors.PanicClose(r.Body)
	content, e := io.ReadAll(r.Body)

	if e != nil {
		return e
	}

	if r.StatusCode >= http.StatusBadRequest {
		return unexpected.Format(
			"%s %s: %d %s",
			method,
			path,
			r.StatusCode,
			content,
		)
	}

	if result == nil || len(content) == 0 {
		return nil
	}

	return json.Unmarshal(content, result)
}
