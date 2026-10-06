package band

import (
	"bytes"
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/digest"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"net/http"
)

func (c *Client) call(
	action string,
	resource string,
	selectors string,
	inner string,
) ([]byte, error) {
	envelope := c.envelope(action, resource, selectors, inner)
	address := join.Empty(c.base, constant.Path)
	first, e := http.NewRequest(
		http.MethodPost,
		address,
		bytes.NewBufferString(envelope),
	)

	if e != nil {
		return nil, e
	}

	first.Header.Set("Content-Type", constant.ContentType)
	challenge, f := c.client.Do(first)

	if f != nil {
		return nil, f
	}

	if challenge.StatusCode == http.StatusOK {
		return read(challenge)
	}

	if g := drain(challenge); g != nil {
		return nil, g
	}

	if challenge.StatusCode != http.StatusUnauthorized {
		return nil, unexpected.Format(
			"band %s: expected digest challenge, got %s",
			resource,
			challenge.Status,
		)
	}

	second, g := http.NewRequest(
		http.MethodPost,
		address,
		bytes.NewBufferString(envelope),
	)

	if g != nil {
		return nil, g
	}

	second.Header.Set("Content-Type", constant.ContentType)
	second.Header.Set(
		"Authorization",
		digest.Authorization(
			challenge.Header.Get("WWW-Authenticate"),
			http.MethodPost,
			constant.Path,
			c.user,
			c.password,
		),
	)
	r, h := c.client.Do(second)

	if h != nil {
		return nil, h
	}

	body, i := io.ReadAll(r.Body)

	if j := r.Body.Close(); j != nil {
		return nil, j
	}

	if i != nil {
		return nil, i
	}

	if r.StatusCode >= http.StatusBadRequest {
		var fault response.Fault

		if xml.Unmarshal(body, &fault) == nil && fault.Text != "" {
			return nil, unexpected.Format(
				"band %s: fault %s",
				resource,
				fault.Text,
			)
		}

		return nil, unexpected.Format("band %s: %s", resource, r.Status)
	}

	return body, nil
}
