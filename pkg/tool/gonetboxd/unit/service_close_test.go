package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func TestServiceClosesTheClientItBuilt(t *testing.T) {
	c := service_tester.NewSource(t)
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gonetboxd.db")
	o.Source = func(prometheus.ObserverVec) face.Client { return c }
	service.New(
		o,
		logger.New(t.Context()),
		metric.New().Upstream(constant.Identity),
	).Close()
	assert.Integer(t, 1, c.Closed)
}
