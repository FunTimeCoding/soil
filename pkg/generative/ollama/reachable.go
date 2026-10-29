package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/ollama/refusal"
)

func (c *Client) Reachable(x context.Context) error {
	return refusal.Classify(c.client.Heartbeat(x))
}
