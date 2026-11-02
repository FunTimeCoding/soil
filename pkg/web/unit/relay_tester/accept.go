package relay_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"sync/atomic"
)

func (r *Relay) accept() {
	for {
		c, e := r.listener.Accept()

		if e != nil {
			return
		}

		u, f := net.Dial(systemConstant.Transmission, r.target)
		errors.PanicOnError(f)
		frozen := &atomic.Bool{}
		r.mutex.Lock()
		r.frozen = append(r.frozen, frozen)
		r.open = append(r.open, c, u)
		r.mutex.Unlock()
		go forward(c, u, frozen)
		go forward(u, c, frozen)
	}
}
