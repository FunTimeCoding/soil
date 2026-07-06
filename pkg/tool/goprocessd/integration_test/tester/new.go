package tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
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
	procfilePath := filepath.Join(directory, "Procfile")
	envrcPath := filepath.Join(directory, ".envrc")
	socketPath := fmt.Sprintf(
		"/tmp/goprocessd-test-%d.sock",
		time.Now().UnixNano(),
	)
	errors.PanicOnError(
		os.WriteFile(procfilePath, []byte(procfileContent), 0644),
	)
	errors.PanicOnError(
		os.WriteFile(envrcPath, []byte(envrcContent), 0755),
	)
	entries, e := procfile.Parse(procfilePath)
	errors.PanicOnError(e)
	env := environment.New(os.Environ())
	errors.PanicOnError(env.Load(envrcPath))
	s := server.New(
		entries,
		env,
		procfilePath,
		envrcPath,
		socketPath,
	)
	go func() { errors.PanicOnError(s.Run()) }()
	time.Sleep(100 * time.Millisecond)
	t.Cleanup(
		func() {
			errors.LogOnError(os.Remove(socketPath))
		},
	)

	return &Tester{
		Server:       s,
		Directory:    directory,
		SocketPath:   socketPath,
		ProcfilePath: procfilePath,
		EnvrcPath:    envrcPath,
	}
}
