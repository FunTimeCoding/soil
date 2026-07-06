package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func (c *Client) send(r *record.Record) {
	body := map[string]any{
		"tool":    r.Tool,
		"surface": r.Surface,
		"actor":   r.Actor,
		"outcome": r.Outcome,
		"kind":    r.Kind,
	}

	if len(r.Detail) > 0 {
		body["detail"] = r.Detail
	}

	encoded, e := json.Marshal(body)

	if e != nil {
		return
	}

	response, e := c.client.Post(
		fmt.Sprintf("%s/api/events", c.base),
		"application/json",
		bytes.NewReader(encoded),
	)

	if e != nil {
		return
	}

	if f := response.Body.Close(); f != nil {
		return
	}
}
