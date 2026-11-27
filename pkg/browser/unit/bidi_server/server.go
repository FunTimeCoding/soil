package bidi_server

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"net/http/httptest"
	"sync"
	"testing"
)

type Server struct {
	t          *testing.T
	server     *httptest.Server
	mutex      sync.Mutex
	answers    map[string]*message.Reply
	connection *websocket.Conn
	received   []*message.Request
	deleted    bool
}
