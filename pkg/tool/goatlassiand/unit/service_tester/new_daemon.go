package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"testing"
)

func NewDaemon(
	t *testing.T,
	o *option.Atlassian,
	r face.Reporter,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
