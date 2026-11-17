package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	argocdMock "github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"
	usageMock "github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"testing"
)

type Tester struct {
	Prometheus *prometheusMock.Client
	Argocd     *argocdMock.Client
	Usage      *usageMock.Client
	Reporter   *memory.Memory
	board      *board.Board
	litePath   string
	service    *service.Service
	daemon     *daemon.Daemon
	t          *testing.T
}
