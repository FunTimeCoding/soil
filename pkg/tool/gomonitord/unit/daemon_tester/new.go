package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gomonitord.db")
	result := daemon.New(o, logger.New(t.Context()))
	t.Cleanup(result.Close)

	return result
}
