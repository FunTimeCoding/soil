package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
)

func (s *Server) Send(
	target string,
	values url.Values,
) (int, string) {
	s.t.Helper()
	v := s.Served()
	var q *http.Request

	if values == nil {
		q = web.NewGet(v.Address(target))
	} else {
		q = web.NewPost(v.Address(target), values.Encode())
		q.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	}

	q.Header.Set(webConstant.ExtendedRequest, "true")
	r, e := v.Client().Do(q)
	assert.FatalOnError(s.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(s.t, f)

	return r.StatusCode, string(b)
}
