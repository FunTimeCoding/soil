package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"testing"
)

func NewDaemon(
	t *testing.T,
	o *option.Log,
	r face.Reporter,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(o, logger.New(t.Context()), r, metric.New())
	t.Cleanup(result.Close)

	return result
}
