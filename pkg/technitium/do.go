package technitium

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/envelope"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"io"
	"net/http"
)

func (c *Client) do(path string) (json.RawMessage, error) {
	r, e := http.NewRequest(
		http.MethodGet,
		join.Empty(c.base, path),
		nil,
	)

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
		return nil, fmt.Errorf(
			"technitium %s: %d (body unreadable: %w)",
			path,
			result.StatusCode,
			g,
		)
	}

	if result.StatusCode >= http.StatusBadRequest {
		return nil, detail_error.New(string(b), result.Status)
	}

	var v envelope.Envelope

	if h := json.Unmarshal(b, &v); h != nil {
		return nil, h
	}

	if v.Status != "ok" {
		return nil, fmt.Errorf("technitium: %s", v.Message)
	}

	return v.Payload, nil
}
