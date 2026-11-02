package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAMalformedNumberIsNotReported(t *testing.T) {
	r := memory.New()
	d, _ := daemon_tester.New(t, r, constant.FixtureInstance)
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
