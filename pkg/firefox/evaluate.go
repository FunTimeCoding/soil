package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/evaluation"
)

func (c *Client) Evaluate(
	identifier int,
	expression string,
) (*evaluation.Evaluation, error) {
	r, e := c.send(
		"evaluate",
		map[string]any{"tab_id": identifier, "expression": expression},
	)

	if e != nil {
		return evaluation.Stub(), fmt.Errorf("evaluate: %w", e)
	}

	var result *evaluation.Evaluation

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return evaluation.Stub(), fmt.Errorf("evaluate: %w", f)
	}

	return result, nil
}
