package base

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"net/http/httptest"
	"testing"
)

type Stack struct {
	T       *testing.T
	Client  *chromium.Client
	Site    *httptest.Server
	Port    int
	process *process.Process
}
