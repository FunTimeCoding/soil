package relay_tester

import (
	"net"
	"sync/atomic"
)

func forward(
	from net.Conn,
	to net.Conn,
	frozen *atomic.Bool,
) {
	b := make([]byte, 32*1024)

	for {
		n, e := from.Read(b)

		if e != nil {
			return
		}

		if frozen.Load() {
			continue
		}

		if _, f := to.Write(b[:n]); f != nil {
			return
		}
	}
}
