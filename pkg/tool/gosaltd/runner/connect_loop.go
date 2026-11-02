package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"time"
)

func (r *Runner) connectLoop(x context.Context) bool {
	if e := peer.Await(
		x,
		r.logger,
		constant.SaltPeer,
		r.salt.Address(),
		r.salt.Login,
	); e != nil {
		if x.Err() != nil {
			return false
		}

		r.logger.Plain("%v", e)

		for !r.connect(x, r.salt) {
			select {
			case <-x.Done():
				return false
			case <-time.After(constant.ConnectRetryInterval):
			}
		}
	}

	close(r.connected)

	return true
}
