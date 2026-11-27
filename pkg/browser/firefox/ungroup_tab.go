package firefox

import "fmt"

func (c *Client) UngroupTab(identifier int) error {
	_, e := c.send("ungroup_tab", map[string]int{"tab_id": identifier})

	if e != nil {
		return fmt.Errorf("ungroup tab: %w", e)
	}

	return nil
}
