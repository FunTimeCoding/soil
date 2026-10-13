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
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			result := evaluation.New()
			result.Identifier = identifier
			result.Value = json.RawMessage(fmt.Sprintf("%q", expression))

			return result, nil
		}
	}

	return evaluation.Stub(), fmt.Errorf("tab %d not found", identifier)
}
