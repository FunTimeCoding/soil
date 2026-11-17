package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/report_page"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) Delete(name string) int {
	s.t.Helper()
	v := s.Served()
	q, e := http.NewRequestWithContext(
		s.t.Context(),
		http.MethodPost,
		v.Address(webConstant.RootPath),
		nil,
	)
	assert.FatalOnError(s.t, e)
	q.URL.Opaque = join.Empty(report_page.Path(name), constant.DeleteSuffix)
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, f := v.Halting().Do(q)
	assert.FatalOnError(s.t, f)
	defer errors.LogClose(r.Body)

	return r.StatusCode
}
