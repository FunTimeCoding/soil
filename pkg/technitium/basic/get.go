package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unreadable_body"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/envelope"
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"net/http"
)

func (c *Client) Get(path string) (json.RawMessage, error) {
	r, e := http.NewRequest(http.MethodGet, join.Empty(c.base, path), nil)

	if e != nil {
		return nil, e
	}

	web.Bearer(r, c.token)
	result, f := c.client.Do(r)

	if f != nil {
		return nil, f
	}

	b, g := io.ReadAll(result.Body)

	if h := result.Body.Close(); h != nil {
		return nil, h
	}

	if g != nil {
		return nil, unreadable_body.New(
			g,
			"technitium %s: %d",
			path,
			result.StatusCode,
		)
	}

	if result.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(b, result.Status)
	}

	var v envelope.Envelope

	if h := json.Unmarshal(b, &v); h != nil {
		return nil, fmt.Errorf("decode envelope: %w", h)
	}

	if v.Status != "ok" {
		return nil, fmt.Errorf("technitium: %s", v.Message)
	}

	return v.Payload, nil
}
