package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) View(identifier int) (*view.View, error) {
	for _, v := range c.views {
		if v.Identifier == identifier {
			return v, nil
		}
	}

	return view.Stub(), fmt.Errorf("view %d not found", identifier)
}
