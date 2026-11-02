package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func RedialService(t *testing.T) (*service.Service, *Shells) {
	t.Helper()
	c := mock_client.New()
	t.Cleanup(c.Close)
	shells := newShells(nil)

	return newServiceWithShells(t, c, shells, memory.New()), shells
}
