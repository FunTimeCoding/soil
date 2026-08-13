package firefox

import "fmt"

func (c *Client) NavigateBack(identifier int) error {
	_, e := c.send("navigate_back", map[string]int{"tab_id": identifier})

	if e != nil {
		return fmt.Errorf("navigate back: %w", e)
	}

	return nil
}
