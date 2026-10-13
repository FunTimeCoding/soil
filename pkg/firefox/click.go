package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (c *Client) Click(
	identifier int,
	selector string,
) (*element.Element, error) {
	r, e := c.send(
		"click",
		map[string]any{"tab_id": identifier, "selector": selector},
	)

	if e != nil {
		return element.Stub(), fmt.Errorf("click: %w", e)
	}

	var result *element.Element

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return element.Stub(), fmt.Errorf("click: %w", f)
	}

	return result, nil
}
