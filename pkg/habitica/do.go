package habitica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/habitica/error_payload"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"io"
	"net/http"
)

func (c *Client) do(
	method string,
	path string,
	body any,
) (*http.Response, error) {
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(notation.Marshal(body))
	}

	r, e := http.NewRequest(method, join.Empty(c.base, path), reader)

	if e != nil {
		return nil, e
	}

	r.Header.Set(constant.UserHeader, c.userIdentifier)
	r.Header.Set(constant.TokenHeader, c.token)
	r.Header.Set(web.ContentType, web.Object)
	result, f := c.client.Do(r)

	if f != nil {
		return nil, f
	}

	if result.StatusCode >= http.StatusBadRequest {
		b, g := io.ReadAll(result.Body)

		if h := result.Body.Close(); h != nil {
			return nil, h
		}

		if g != nil {
			return nil, fmt.Errorf(
				"habitica %s %s: %d (body unreadable: %w)",
				method,
				path,
				result.StatusCode,
				g,
			)
		}

		var o error_payload.Payload

		if json.Unmarshal(b, &o) == nil && o.Message != "" {
			return nil, detail_error.New(
				o.Message,
				result.Status,
			)
		}

		return nil, fmt.Errorf("%s", result.Status)
	}

	return result, nil
}
