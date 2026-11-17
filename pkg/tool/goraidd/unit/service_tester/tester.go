package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client/mock_client"
	"testing"
	"time"
)

type Tester struct {
	Parser   *mock_client.Client
	Reporter *memory.Memory
	LogCache string
	Output   string
	Elite    string
	litePath string
	clock    func() time.Time
	service  *service.Service
	daemon   *daemon.Daemon
	t        *testing.T
}
