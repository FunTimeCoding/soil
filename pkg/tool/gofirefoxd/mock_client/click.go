package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (c *Client) Click(
	identifier int,
	selector string,
) (*element.Element, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			result := element.New()
			result.Selector = selector
			result.Tag = "button"

			return result, nil
		}
	}

	return element.Stub(), fmt.Errorf("tab %d not found", identifier)
}
