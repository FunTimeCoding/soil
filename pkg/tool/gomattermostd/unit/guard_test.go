package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.Mount(
		t,
		service_tester.New(t).Daemon(),
		memory.New(),
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
