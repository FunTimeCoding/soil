package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
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
		func(string) (face.ProxmoxClient, error) { return c, nil },
		inventory.New[instance.Instance](
			instance.Instance{Name: name, Host: "mock"},
		),
		time.Minute,
		y,
		logger.New(context.Background()),
		reporter.NewOptional(constant.Identity.Name()),
	), y
}
