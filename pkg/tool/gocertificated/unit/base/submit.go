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
) string {
	s.t.Helper()
	v := s.Served()
	r, e := http.NewRequest(
		http.MethodPost,
		v.Address(path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(s.t, e)
	r.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	web.Bearer(r, generativeConstant.ModelContextTestToken)
	result, e := v.Client().Do(r)
	assert.FatalOnError(s.t, e)

	return web.ReadString(result)
}
