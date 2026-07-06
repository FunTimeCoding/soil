package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func (c *Client) exchange(
	method string,
	path string,
	body any,
) ([]byte, int, error) {
	var reader io.Reader

	if body != nil {
		b, e := json.Marshal(body)

		if e != nil {
			return nil, 0, e
		}

		reader = bytes.NewReader(b)
	}

	r, e := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", c.base, path),
		reader,
	)

	if e != nil {
		return nil, 0, e
	}

	r.Header.Set(web.Accept, web.Object)

	if body != nil {
		r.Header.Set(web.ContentType, web.Object)
	}

	if c.token != "" {
		r.Header.Set(constant.TokenHeader, c.token)
	}

	s, f := c.client.Do(r)

	if f != nil {
		return nil, 0, f
	}

	b, g := io.ReadAll(s.Body)
	h := s.Body.Close()

	if g != nil {
		return nil, 0, g
	}

	if h != nil {
		return nil, 0, h
	}

	return b, s.StatusCode, nil
}
