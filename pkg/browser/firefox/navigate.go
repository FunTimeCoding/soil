package firefox

import "fmt"

func (c *Client) Navigate(
	identifier int,
	l string,
) error {
	_, e := c.send("navigate", map[string]any{"tab_id": identifier, "url": l})

	if e != nil {
		return fmt.Errorf("navigate: %w", e)
	}

	return nil
}
