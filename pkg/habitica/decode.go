package habitica

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/habitica/envelope"
	"io"
	"net/http"
)

func (c *Client) decode(
	r *http.Response,
	result any,
) error {
	b, e := io.ReadAll(r.Body)

	if f := r.Body.Close(); f != nil {
		return f
	}

	if e != nil {
		return e
	}

	var envelope *envelope.Envelope

	if f := json.Unmarshal(b, &envelope); f != nil {
		return f
	}

	return json.Unmarshal(envelope.Payload, result)
}
