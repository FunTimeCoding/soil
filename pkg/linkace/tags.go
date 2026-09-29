package linkace

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	var result []*tag.Tag
	p := 1

	for {
		r, e := c.TagsPage(p)

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
