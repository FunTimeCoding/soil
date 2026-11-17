package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/unit/daemon_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.Mount(
		t,
		daemon_tester.New(t, "alfa: sleep 60\n"),
		memory.New(),
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
	v.VerifyGuarded(t, "/api/processes")
	v.VerifyGuarded(t, "/api/processes/alfa/log")
}
