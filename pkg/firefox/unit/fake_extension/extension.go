package fake_extension

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
	"sync"
)

type Extension struct {
	mutex      sync.Mutex
	connection *websocket.Conn
	replies    map[string]*message.Reply
}
