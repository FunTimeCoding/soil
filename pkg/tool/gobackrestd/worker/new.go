package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/types/summary"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	latest func(context.Context) ([]*summary.Summary, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	result := &Worker{
		latest: latest,
		status: gauge(
			y,
			constant.StatusMetric,
			constant.StatusHelp,
			constant.StanzaLabel,
		),
		completion: gauge(
			y,
			constant.CompletionMetric,
			constant.CompletionHelp,
			constant.StanzaLabel,
			constant.TypeLabel,
		),
		size: gauge(
			y,
			constant.SizeMetric,
			constant.SizeHelp,
			constant.StanzaLabel,
			constant.TypeLabel,
		),
		repository: gauge(
			y,
			constant.RepositorySizeMetric,
			constant.RepositorySizeHelp,
			constant.StanzaLabel,
			constant.TypeLabel,
		),
		damage: gauge(
			y,
			constant.ErrorMetric,
			constant.ErrorHelp,
			constant.StanzaLabel,
			constant.TypeLabel,
		),
	}
	result.clock = worker.New(result.tick, l, r, interval)

	return result
}
