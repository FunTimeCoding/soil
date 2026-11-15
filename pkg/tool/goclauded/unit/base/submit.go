package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) Submit(
	path string,
	values url.Values,
	extended bool,
) string {
	s.t.Helper()
	q, e := http.NewRequest(
		http.MethodPost,
		s.Address(path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(s.t, e)
	q.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	web.Bearer(q, generativeConstant.ModelContextTestToken)

	if extended {
		q.Header.Set(webConstant.ExtendedRequest, "true")
	}

	r, f := s.Server.Client().Do(q)
	assert.FatalOnError(s.t, f)

	return web.ReadString(r)
}
