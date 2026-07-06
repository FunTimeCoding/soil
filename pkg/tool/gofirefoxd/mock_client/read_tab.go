package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/content"
)

func (c *Client) ReadTab(
	identifier int,
	_ bool,
) (*content.Content, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			result := content.New()
			result.Identifier = t.Identifier
			result.Locator = t.Locator
			result.Title = t.Title

			return result, nil
		}
	}

	return content.Stub(), fmt.Errorf("tab %d not found", identifier)
}
