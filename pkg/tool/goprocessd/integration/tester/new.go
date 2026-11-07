package tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func New(
	t *testing.T,
	procfileContent string,
	envrcContent string,
) *Tester {
	t.Helper()
	directory := t.TempDir()
	o := option.New()
	o.ProcfilePath = filepath.Join(directory, "Procfile")
	o.EnvrcPath = filepath.Join(directory, ".envrc")
	errors.PanicOnError(
		os.WriteFile(o.ProcfilePath, []byte(procfileContent), 0644),
	)
	errors.PanicOnError(os.WriteFile(o.EnvrcPath, []byte(envrcContent), 0755))
	d := daemon.New(o)
	m := http.NewServeMux()
	d.MountLocal(memory.New(), mock_recorder.New(), guard.New(m, nil))
	l := socket.Listen(socket.Path(o.ProcfilePath))
	h := &http.Server{Handler: m}
	go func() { errors.PanicOnUnclean(h.Serve(l)) }()
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	time.Sleep(100 * time.Millisecond)
	c := client.NewLocal(o.ProcfilePath)
	t.Cleanup(
		func() {
			c.Close()
			errors.LogOnError(h.Close())

			for _, w := range workers {
				w.Close()
			}

			d.Close()
		},
	)

	return &Tester{
		Client:       c,
		Directory:    directory,
		ProcfilePath: o.ProcfilePath,
		EnvrcPath:    o.EnvrcPath,
	}
}
