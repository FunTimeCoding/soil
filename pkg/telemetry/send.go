package telemetry

import (
	"context"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
)

func (c *Client) send(r *record.Record) {
	body := client.PostEventJSONRequestBody{
		Tool:    r.Tool,
		Surface: r.Surface,
		Actor:   r.Actor,
		Outcome: r.Outcome,
		Kind:    r.Kind,
	}

	if len(r.Detail) > 0 {
		body.Detail = &r.Detail
	}

	if _, e := c.client.PostEvent(context.Background(), body); e != nil {
		return
	}
}
