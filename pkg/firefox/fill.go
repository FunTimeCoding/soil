package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (c *Client) Fill(
	identifier int,
	selector string,
	value string,
	direct bool,
) (*element.Element, error) {
	r, e := c.send(
		"fill",
		map[string]any{
			"tab_id":   identifier,
			"selector": selector,
			"value":    value,
			"direct":   direct,
		},
	)

	if e != nil {
		return element.Stub(), fmt.Errorf("fill: %w", e)
	}

	var result *element.Element

	if f := json.Unmarshal(r.Result, &result); f != nil {
		return element.Stub(), fmt.Errorf("fill: %w", f)
	}

	return result, nil
}
