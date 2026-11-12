package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/unit/upstream_tester/answer"
	"net/http/httptest"
	"sync"
	"testing"
)

type Upstream struct {
	t       *testing.T
	server  *httptest.Server
	mutex   sync.Mutex
	answers map[string]*answer.Answer
}
