package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"time"
)

func (r *Runner) keySyncLoop() {
	select {
	case <-r.connected:
	case <-r.provision.Done():
		return
	}

	ticker := time.NewTicker(constant.KeySyncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-r.provision.Done():
			return
		case <-ticker.C:
			r.recovery.Run(r.keySync)
		}
	}
}
