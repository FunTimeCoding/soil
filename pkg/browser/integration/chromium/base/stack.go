package base

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"net/http/httptest"
	"testing"
)

type Stack struct {
	T       *testing.T
	Client  *chromium.Client
	Site    *httptest.Server
	Port    int
	profile string
	process *process.Process
	pool    *pool.Pool
}
