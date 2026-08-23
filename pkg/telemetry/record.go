package telemetry

import "github.com/funtimecoding/soil/pkg/telemetry/record"

func (c *Client) Record(r *record.Record) {
	c.group.Add(1)
	go func() {
		defer c.group.Done()
		c.send(r)
	}()
}
