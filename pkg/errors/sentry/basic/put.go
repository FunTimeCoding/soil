package basic

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"io"
	"net/http"
)

func (c *Client) Put(
	path string,
	body any,
) ([]byte, error) {
	b, e := json.Marshal(body)

	if e != nil {
		return nil, e
	}

	r, f := http.NewRequest(
		http.MethodPut,
		locator.New(c.host).Base(constant.Base).Path(path).Trail().String(),
		bytes.NewReader(b),
	)

	if f != nil {
		return nil, f
	}

	web.Bearer(r, c.token)
	r.Header.Add(webConstant.ContentType, webConstant.Object)
	r.Header.Add(webConstant.Accept, webConstant.Object)
	s, g := http.DefaultClient.Do(r)

	if g != nil {
		return nil, g
	}

	result, h := io.ReadAll(s.Body)
	i := s.Body.Close()

	if h != nil {
		return nil, h
	}

	if i != nil {
		return nil, i
	}

	if s.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(result, s.Status)
	}

	return result, nil
}
