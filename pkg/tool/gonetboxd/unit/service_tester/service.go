package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func Service(
	t *testing.T,
	c face.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gonetboxd.db")
	o.Source = func(prometheus.ObserverVec) face.Client { return c }
	result := service.New(
		o,
		logger.New(t.Context()),
		metric.New().Upstream(constant.Identity),
	)
	t.Cleanup(result.Close)

	return result
}
