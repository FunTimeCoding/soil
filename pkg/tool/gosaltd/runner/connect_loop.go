package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"time"
)

func (r *Runner) connectLoop() bool {
	for r.salt == nil {
		r.recovery.Run(r.connect)

		if r.salt != nil {
			close(r.connected)

			return true
		}

		select {
		case <-r.provision.Done():
			return false
		case <-time.After(constant.ConnectRetryInterval):
		}
	}

	return true
}
