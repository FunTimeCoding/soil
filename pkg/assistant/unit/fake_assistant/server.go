package fake_assistant

import (
	"github.com/gorilla/websocket"
	"net/http/httptest"
	"sync"
)

type Server struct {
	mutex        sync.Mutex
	server       *httptest.Server
	refusal      int
	invalid      bool
	subscription int
	socket       *websocket.Conn
	identifier   uint64
}
