package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetValues(
	path string,
	query url.Values,
) ([]byte, error) {
	l := locator.New(c.host).Base(constant.Base).Path(path).Trail()

	for k, list := range query {
		for _, v := range list {
			l.Add(k, v)
		}
	}

	r, e := http.NewRequest(http.MethodGet, l.String(), nil)

	if e != nil {
		return nil, e
	}

	web.Bearer(r, c.token)
	s, f := http.DefaultClient.Do(r)

	if f != nil {
		return nil, f
	}

	result, g := io.ReadAll(s.Body)
	h := s.Body.Close()

	if g != nil {
		return nil, g
	}

	if h != nil {
		return nil, h
	}

	if s.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(result, s.Status)
	}

	return result, nil
}
