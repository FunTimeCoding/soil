package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/daemon_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.Mount(
		t,
		daemon_tester.New(t, mock_client.New(t)),
		memory.New(),
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
