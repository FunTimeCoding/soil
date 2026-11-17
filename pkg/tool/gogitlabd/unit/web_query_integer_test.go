package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMalformedNumberIsNotReported(t *testing.T) {
	s := base.New(t)
	r := s.Reporter
	d := s.Daemon()
	m := http.NewServeMux()
	d.Mount(
		r,
		mock_recorder.New(),
		guard.New(m, []string{generativeConstant.ModelContextTestToken}),
	)
	w := httptest.NewRecorder()
	d.Recovery(r)(m).ServeHTTP(
		w,
		httptest.NewRequest(
			http.MethodGet,
			"/pipeline?instance=alfa&project=kilo&pipeline=1",
			nil,
		),
	)
	assert.True(t, w.Code != http.StatusOK)
	assert.Count(t, 0, r.Events())
}
