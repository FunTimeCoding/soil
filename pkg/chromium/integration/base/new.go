package base

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/system/run"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Stack {
	t.Helper()
	port := freePort(t)
	result := &Stack{
		T:    t,
		Site: httptest.NewServer(http.HandlerFunc(serve)),
		Port: port,
		process: run.New().Open(
			constant.BravePath,
			"--headless=new",
			fmt.Sprintf("--remote-debugging-port=%d", port),
			fmt.Sprintf("--user-data-dir=%s", t.TempDir()),
			"--no-first-run",
			"--no-default-browser-check",
			"--disable-gpu",
			"about:blank",
		),
	}
	t.Cleanup(result.Close)
	waitReady(t, port)
	result.Client = chromium.New("localhost", port)

	return result
}
