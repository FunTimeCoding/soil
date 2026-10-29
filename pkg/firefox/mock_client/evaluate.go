package mock_client

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/evaluation"
)

func (c *Client) Evaluate(
	identifier int,
	expression string,
) (*evaluation.Evaluation, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return evaluation.Stub(), e
	}

	if _, _, f := c.findTab(identifier); f != nil {
		return evaluation.Stub(), f
	}

	result := evaluation.New()
	result.Identifier = identifier
	result.Value = json.RawMessage(fmt.Sprintf("%q", expression))

	return result, nil
}
