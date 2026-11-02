package runner

import "time"

func (r *Runner) keySyncLoop() {
	defer r.keySyncing.Done()

	select {
	case <-r.connected:
	case <-r.provision.Done():
		return
	}

	ticker := time.NewTicker(r.keySyncInterval)
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
