package firefox

import "fmt"

func (c *Client) UpdateGroup(
	groupIdentifier int,
	title string,
	color string,
	collapsed *bool,
) error {
	parameters := map[string]any{"group_id": groupIdentifier}

	if title != "" {
		parameters["title"] = title
	}

	if color != "" {
		parameters["color"] = color
	}

	if collapsed != nil {
		parameters["collapsed"] = *collapsed
	}

	_, e := c.send("update_group", parameters)

	if e != nil {
		return fmt.Errorf("update group: %w", e)
	}

	return nil
}
