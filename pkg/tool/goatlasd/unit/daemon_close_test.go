package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestCloseClosesEveryPeer(t *testing.T) {
	v := service_tester.New(t)
	d := daemon.New(
		v.Option,
		logger.New(context.Background()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	d.Close()
	assert.Integer(t, 1, v.Netbox.Closed)
	assert.Integer(t, 1, v.Process.Closed)
	assert.Integer(t, 1, v.Lease.Closed)
}
