package relay_tester

import (
	"net"
	"sync"
	"sync/atomic"
)

type Relay struct {
	listener net.Listener
	target   string
	mutex    sync.Mutex
	frozen   []*atomic.Bool
	open     []net.Conn
}
