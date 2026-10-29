package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"time"
)

func (r *Runner) connectLoop(x context.Context) bool {
	c := r.saltConnector()
	r.keep(c)

	if e := peer.Await(
		x,
		r.logger,
		constant.SaltPeer,
		c.Address(),
		c.Login,
	); e != nil {
		if x.Err() != nil {
			return false
		}

		r.logger.Plain("%v", e)

		for !r.connect(x, c) {
			select {
			case <-x.Done():
				return false
			case <-time.After(constant.ConnectRetryInterval):
			}
		}
	}

	r.salt = c
	close(r.connected)

	return true
}
