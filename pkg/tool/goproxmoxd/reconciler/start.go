package reconciler

import "time"

func (r *Reconciler) Start() {
	go func() {
		t := time.NewTicker(r.interval)
		defer t.Stop()
		r.recovery.Run(r.Reconcile)

		for {
			select {
			case <-r.context.Done():
				return
			case <-t.C:
				r.recovery.Run(r.Reconcile)
			}
		}
	}()
}
