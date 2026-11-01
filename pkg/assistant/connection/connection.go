package connection

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	sync.RWMutex
	host           string
	token          string
	subscriber     Subscriber
	connection     *websocket.Conn
	subscribers    map[uint64]Subscriber
	lastIdentifier uint64
}
