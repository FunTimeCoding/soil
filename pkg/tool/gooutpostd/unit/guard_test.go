package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/daemon"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.Mount(t, daemon.New(), memory.New())
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/services")
	v.VerifyGuarded(t, "/api/host")
}
