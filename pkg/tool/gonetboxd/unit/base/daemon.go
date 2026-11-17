package base

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func Daemon(
	t *testing.T,
	c face.NetBoxClient,
) *daemon.Daemon {
	t.Helper()

	return service_tester.New(t, service_tester.WithClient(c)).Daemon()
}
