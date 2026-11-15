package proxmox_server

import (
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server/answer"
	"net/http/httptest"
	"sync"
	"testing"
)

type Server struct {
	t       *testing.T
	server  *httptest.Server
	mutex   sync.Mutex
	answers map[string]*answer.Answer
}
