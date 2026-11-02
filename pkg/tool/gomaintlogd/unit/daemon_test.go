package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"go.uber.org/goleak"
	"path/filepath"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "maintenance.db")
	d := daemon.New(o, logger.New(t.Context()), memory.New())
	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
