package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/metric"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTheExporterMountsOpenAtTheMetricsPath(t *testing.T) {
	m := http.NewServeMux()
	metric.New().Mount(guard.New(m, nil))
	r := httptest.NewRecorder()
	m.ServeHTTP(
		r,
		httptest.NewRequest(http.MethodGet, webConstant.MetricsPath, nil),
	)
	assert.Integer(t, http.StatusOK, r.Code)
}
