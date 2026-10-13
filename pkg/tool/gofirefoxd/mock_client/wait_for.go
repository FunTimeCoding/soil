package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) WaitFor(
	identifier int,
	_ string,
	_ int,
) (*tab.Tab, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			return t, nil
		}
	}

	return tab.Stub(), fmt.Errorf("tab %d not found", identifier)
}
