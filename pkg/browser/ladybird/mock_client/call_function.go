package mock_client

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/types/call"
)

func (c *Client) CallFunction(
	context string,
	declaration string,
	arguments ...any,
) (json.RawMessage, error) {
	c.mutex.Lock()

	if _, e := c.find(context); e != nil {
		c.mutex.Unlock()

		return nil, e
	}

	c.calls = append(c.calls, call.New(context, declaration, arguments))
	refusal, refused := c.refusals[declaration]
	result, answered := c.results[declaration]
	events := c.navigations[declaration]
	listen := c.event
	c.mutex.Unlock()

	if listen != nil {
		for _, r := range events {
			listen(r)
		}
	}

	if refused {
		return nil, refusal
	}

	if answered {
		return result, nil
	}

	return json.RawMessage(`{"type":"undefined"}`), nil
}
