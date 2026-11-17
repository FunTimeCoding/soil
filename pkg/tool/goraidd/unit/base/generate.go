package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) Generate(names ...string) (int, string) {
	s.t.Helper()
	v := s.Served()
	q, e := http.NewRequestWithContext(
		s.t.Context(),
		http.MethodPost,
		v.Address(constant.GeneratePath),
		strings.NewReader(url.Values{constant.FileNameField: names}.Encode()),
	)
	assert.FatalOnError(s.t, e)
	q.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, f := v.Halting().Do(q)
	assert.FatalOnError(s.t, f)
	defer errors.LogClose(r.Body)
	b, g := io.ReadAll(r.Body)
	assert.FatalOnError(s.t, g)

	return r.StatusCode, string(b)
}
