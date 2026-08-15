package janitor

import "time"

func (j *Janitor) Start() {
	go func() {
		t := time.NewTicker(j.interval)
		defer t.Stop()
		j.recovery.Run(j.sweep)

		for {
			select {
			case <-t.C:
				j.recovery.Run(j.sweep)
			case <-j.stop:
				return
			}
		}
	}()
}
