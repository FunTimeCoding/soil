package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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
		if v.Identifier != identifier {
			continue
		}

		if !all && strings.Count(v.Text, old) > 1 {
			return view.Stub(), conflict.Format(
				"edit view: old_string matches multiple times; use replace_all",
			)
		}

		if all {
			c.views[i].Text = strings.ReplaceAll(v.Text, old, new)
		} else {
			c.views[i].Text = strings.Replace(v.Text, old, new, 1)
		}

		c.views[i].Dirty = true

		return c.views[i], nil
	}

	return view.Stub(), not_found.New("view", identifier)
}
