package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) TagByName(name string) (*tag.Tag, error) {
	all, e := c.Tags()

	if e != nil {
		return nil, e
	}

	for _, t := range all {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, not_found.New(constant.TagSubject, name)
}
