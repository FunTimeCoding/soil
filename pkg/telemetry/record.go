package telemetry

import "github.com/funtimecoding/soil/pkg/telemetry/record"

func (c *Client) Record(r *record.Record) {
	go c.send(r)
}
