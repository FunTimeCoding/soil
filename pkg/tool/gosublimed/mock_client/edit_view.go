package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"strings"
)

func (c *Client) EditView(
	identifier int,
	old string,
	new string,
	all bool,
) (*view.View, error) {
	for i, v := range c.views {
		if v.Identifier == identifier {
			if all {
				c.views[i].Text = strings.ReplaceAll(v.Text, old, new)
			} else {
				c.views[i].Text = strings.Replace(v.Text, old, new, 1)
			}

			c.views[i].Dirty = true

			return c.views[i], nil
		}
	}

	return view.Stub(), fmt.Errorf("view %d not found", identifier)
}
