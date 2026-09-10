package gauge

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/prometheus/client_golang/prometheus"
)

type Gauge struct {
	service  *service.Service
	logger   *logger.Logger
	findings *prometheus.GaugeVec
}
