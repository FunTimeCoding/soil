package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) Page(path string) string {
	s.t.Helper()
	v := s.Served()
	r, e := http.NewRequest(http.MethodGet, v.Address(path), nil)
	assert.FatalOnError(s.t, e)
	web.Bearer(r, generativeConstant.ModelContextTestToken)
	result, f := v.Client().Do(r)
	assert.FatalOnError(s.t, f)
	assert.Integer(s.t, http.StatusOK, result.StatusCode)

	return web.ReadString(result)
}
