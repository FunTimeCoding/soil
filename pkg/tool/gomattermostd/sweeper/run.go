package sweeper

import "time"

func (s *Sweeper) run(t *time.Ticker) {
	defer t.Stop()

	for {
		select {
		case <-s.done:
			return
		case <-t.C:
			s.Sweep()
		}
	}
}
