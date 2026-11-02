package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"
	"go.uber.org/goleak"
	"os"
	"path/filepath"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	directory := t.TempDir()
	o := option.New()
	o.ProcfilePath = filepath.Join(directory, "Procfile")
	o.EnvrcPath = filepath.Join(directory, ".envrc")
	errors.PanicOnError(
		os.WriteFile(
			o.ProcfilePath,
			[]byte("alfa: sleep 60\nbravo: sleep 60\n"),
			0644,
		),
	)
	before := goleak.IgnoreCurrent()
	d := daemon.New(o)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestAProcessDroppedOnReloadLeavesNothingRunning(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "Procfile")
	errors.PanicOnError(
		os.WriteFile(path, []byte("alfa: sleep 60\nbravo: sleep 60\n"), 0644),
	)
	entries, e := procfile.Parse(path)
	errors.PanicOnError(e)
	before := goleak.IgnoreCurrent()
	s := supervisor.New(
		entries,
		environment.New(os.Environ()),
		path,
		filepath.Join(directory, ".envrc"),
	)
	s.Start()
	errors.PanicOnError(os.WriteFile(path, []byte("alfa: sleep 60\n"), 0644))
	errors.PanicOnError(s.ReloadProcfile())
	s.Close()
	goleak.VerifyNone(t, before)
}
