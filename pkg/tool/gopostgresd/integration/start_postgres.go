package integration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"strconv"
	"testing"
	"time"
)

func startPostgres(
	t *testing.T,
	directory string,
) (string, *instance.Instance) {
	t.Helper()
	port := freePort(t)
	container := fmt.Sprintf("gopostgresd-backup-%d", port)
	i := instance.Instance{
		Name:     "alfa",
		Host:     "127.0.0.1",
		Port:     port,
		Database: "alfa",
		User:     "alfa",
		Password: "alfa",
	}
	r := run.New().NoPanic()
	r.Start(
		"podman",
		"run",
		"--detach",
		"--rm",
		"--name",
		container,
		"--env",
		"POSTGRES_USER=alfa",
		"--env",
		"POSTGRES_PASSWORD=alfa",
		"--env",
		"POSTGRES_DB=alfa",
		"--publish",
		fmt.Sprintf("127.0.0.1:%d:%d", port, port),
		"--volume",
		fmt.Sprintf("%s:%s", directory, directory),
		"docker.io/library/postgres:16",
		"-p",
		strconv.Itoa(port),
	)
	assert.FatalOnError(t, r.Error)
	t.Cleanup(
		func() {
			run.New().NoPanic().Start("podman", "rm", "--force", container)
		},
	)
	deadline := time.Now().Add(time.Minute)

	for {
		ready := run.New().NoPanic()
		ready.Start(
			"podman",
			"exec",
			container,
			"pg_isready",
			"--host",
			i.Host,
			"--port",
			strconv.Itoa(port),
		)

		if ready.Error == nil {
			return container, &i
		}

		if time.Now().After(deadline) {
			t.Fatalf("postgres not ready: %v", ready.Error)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
