package base

import (
	"github.com/funtimecoding/soil/pkg/chromium"
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
	process *process.Process
	pool    *pool.Pool
}
