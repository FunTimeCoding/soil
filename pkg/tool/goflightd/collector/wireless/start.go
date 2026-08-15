package wireless

import "time"

func (c *Collector) Start() {
	go func() {
		t := time.NewTicker(c.interval)
		defer t.Stop()
		c.recovery.Run(c.poll)

		for {
			select {
			case <-t.C:
				c.recovery.Run(c.poll)
			case <-c.stop:
				return
			}
		}
	}()
}
