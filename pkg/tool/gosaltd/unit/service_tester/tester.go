package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

type Tester struct {
	Reporter        *memory.Memory
	Registry        *prometheus.Registry
	master          func() face.SaltClient
	keySyncInterval time.Duration
	litePath        string
	service         *service.Service
	runner          *runner.Runner
	t               *testing.T
}
