package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
	"time"
)

type Tester struct {
	Indexer       *mock_client.Client
	Reporter      *memory.Memory
	Path          string
	hiddenTag     string
	referenceRoot string
	clock         func() time.Time
	service       *service.Service
	daemon        *daemon.Daemon
	t             *testing.T
}
