package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func NewPollWorker(
	name string,
	c *mock_client.Client,
) (*worker.Worker, *prometheus.Registry) {
	y := prometheus.NewRegistry()

	return worker.New(
		mock_service.New(name, c, nil),
		time.Minute,
		y,
		logger.New(context.Background()),
		reporter.NewOptional(constant.Identity.Name(), "test"),
	), y
}
