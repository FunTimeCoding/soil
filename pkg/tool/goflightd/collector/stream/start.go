package stream

import "time"

func (c *Collector) Start() {
	go func() {
		for {
			c.recovery.Run(c.capture)

			select {
			case <-c.stop:
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
}
