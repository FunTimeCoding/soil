package firefox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/response"
)

func (c *Client) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	parameters := map[string]any{
		"tab_ids": tabIdentifiers,
	}

	if groupIdentifier > 0 {
		parameters["group_id"] = groupIdentifier
	}

	if title != "" {
		parameters["title"] = title
	}

	if color != "" {
		parameters["color"] = color
	}

	r, e := c.send("group_tabs", parameters)

	if e != nil {
		return 0, fmt.Errorf("group tabs: %w", e)
	}

	var result *response.GroupTabs

	if f := decodeResult(r, &result); f != nil {
		return 0, fmt.Errorf("group tabs: %w", f)
	}

	return result.GroupIdentifier, nil
}
