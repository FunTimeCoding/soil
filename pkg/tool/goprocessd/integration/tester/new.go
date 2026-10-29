package tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
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
	v := service.New(o)
	m := http.NewServeMux()
	goprocessd.MountLocal(v, memory.New(), mock_recorder.New(), m)
	l := socket.Listen(socket.Path(o.ProcfilePath))
	h := &http.Server{Handler: m}
	go func() { errors.PanicOnUnclean(h.Serve(l)) }()
	go func() { errors.PanicOnError(v.Run()) }()
	time.Sleep(100 * time.Millisecond)
	c := client.NewLocal(o.ProcfilePath)
	t.Cleanup(
		func() {
			c.Close()
			errors.LogOnError(h.Close())
		},
	)

	return &Tester{
		Client:       c,
		Directory:    directory,
		ProcfilePath: o.ProcfilePath,
		EnvrcPath:    o.EnvrcPath,
	}
}
