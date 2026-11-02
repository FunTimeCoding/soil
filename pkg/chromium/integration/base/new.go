package base

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Stack {
	t.Helper()
	result := &Stack{
		T:       t,
		Site:    httptest.NewServer(http.HandlerFunc(serve)),
		Port:    freePort(t),
		profile: t.TempDir(),
		pool:    stall_client.New().Build(),
	}
	t.Cleanup(result.Close)
	result.Start()
	result.Client = chromium.New("localhost", result.Port)

	return result
}
