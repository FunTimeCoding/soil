package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"testing"
)

func NewDaemon(
	t *testing.T,
	o *option.Server,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
