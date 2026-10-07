package firefox

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
	"sync"
	"sync/atomic"
)

type Client struct {
	address    string
	connection *websocket.Conn
	mutex      sync.Mutex
	pending    map[int]chan *message.Reply
	identifier atomic.Int64
}
