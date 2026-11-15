package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/gorilla/websocket"
	"net/http/httptest"
	"sync"
	"testing"
)

type Tester struct {
	Client     *mattermost.Client
	server     *httptest.Server
	t          *testing.T
	connection *websocket.Conn
	accepted   int
	refusals   int
	pings      int
	released   int
	ready      chan struct{}
	once       sync.Once
	mutex      sync.Mutex
}
