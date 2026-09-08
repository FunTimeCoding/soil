package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/gorilla/websocket"
	"sync"
	"testing"
)

type Tester struct {
	Client     *mattermost.Client
	t          *testing.T
	connection *websocket.Conn
	accepted   int
	ready      chan struct{}
	once       sync.Once
	mutex      sync.Mutex
}
