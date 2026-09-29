package linkace

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (c *Client) LinksByList(identifier int) ([]*link.Link, error) {
	var result []*link.Link
	p := 1

	for {
		r, e := c.LinksByListPage(identifier, p)

		if e != nil {
			return nil, e
		}

		result = append(result, r.Items...)

		if p >= r.LastPage {
			break
		}

		p++
	}

	return result, nil
}
