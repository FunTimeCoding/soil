package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"testing"
)

type Tester struct {
	server   *base.Server
	Client   *client.ClientWithResponses
	Store    *store.Store
	Daemon   *daemon.Daemon
	Reporter *memory.Memory
	t        *testing.T
}
