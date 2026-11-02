package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"os"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	procfile string,
) *daemon.Daemon {
	t.Helper()
	directory := t.TempDir()
	o := option.New()
	o.ProcfilePath = filepath.Join(directory, "Procfile")
	o.EnvrcPath = filepath.Join(directory, ".envrc")
	errors.PanicOnError(os.WriteFile(o.ProcfilePath, []byte(procfile), 0644))
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
