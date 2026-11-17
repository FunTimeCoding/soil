package base

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) Post(
	path string,
	body []byte,
) int {
	s.t.Helper()
	v := s.Served()
	q, e := http.NewRequestWithContext(
		s.t.Context(),
		http.MethodPost,
		v.Address(webConstant.RootPath),
		bytes.NewReader(body),
	)
	assert.FatalOnError(s.t, e)
	q.URL.Opaque = path
	q.Header.Set(webConstant.ContentType, webConstant.Octets)
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, f := v.Halting().Do(q)
	assert.FatalOnError(s.t, f)
	defer errors.LogClose(r.Body)

	return r.StatusCode
}
