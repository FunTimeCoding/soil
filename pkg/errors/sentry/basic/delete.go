package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unreadable_body"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"io"
	"net/http"
)

func (c *Client) Delete(path string) error {
	r, e := http.NewRequest(
		http.MethodDelete,
		locator.New(c.host).Base(constant.Base).Path(path).Trail().String(),
		nil,
	)

	if e != nil {
		return e
	}

	web.Bearer(r, c.token)
	s, f := http.DefaultClient.Do(r)

	if f != nil {
		return f
	}

	result, g := io.ReadAll(s.Body)
	h := s.Body.Close()

	if g != nil {
		return unreadable_body.New(g, "read response body")
	}

	if h != nil {
		return h
	}

	if s.StatusCode >= http.StatusBadRequest {
		return parseDetail(result, s.Status)
	}

	return nil
}
