package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/element"
)

func (c *Client) Fill(
	identifier int,
	selector string,
	value string,
	_ bool,
) (*element.Element, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			result := element.New()
			result.Selector = selector
			result.Tag = "input"
			result.Text = value

			return result, nil
		}
	}

	return element.Stub(), fmt.Errorf("tab %d not found", identifier)
}
