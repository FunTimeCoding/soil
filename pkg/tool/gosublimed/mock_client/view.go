package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) View(identifier int) (*view.View, error) {
	for _, v := range c.views {
		if v.Identifier == identifier {
			return v, nil
		}
	}

	return view.Stub(), not_found.New("view", identifier)
}
