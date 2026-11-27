package ladybird

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"sync"
	"sync/atomic"
)

type Client struct {
	option     *ladybird_option.Option
	pool       *pool.Pool
	driver     *process.Process
	address    string
	session    string
	connection *websocket.Conn
	mutex      sync.Mutex
	pending    map[int]chan *message.Reply
	identifier atomic.Int64
}
