package gauge

import "time"

func (g *Gauge) Start() {
	g.started.Do(
		func() {
			go func() {
				defer close(g.done)
				t := time.NewTicker(g.interval)
				defer t.Stop()
				g.recovery.Run(g.Poll)

				for {
					select {
					case <-g.context.Done():
						return
					case <-t.C:
						g.recovery.Run(g.Poll)
					}
				}
			}()
		},
	)
}
