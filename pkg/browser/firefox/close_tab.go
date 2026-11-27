package firefox

import "fmt"

func (c *Client) CloseTab(identifier int) error {
	_, e := c.send("close_tab", map[string]int{"tab_id": identifier})

	if e != nil {
		return fmt.Errorf("close tab: %w", e)
	}

	return nil
}
