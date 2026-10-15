package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) ChoiceSetByName(n string) (*choice_set.Set, error) {
	result, e := c.ChoiceSets()

	if e != nil {
		return nil, e
	}

	for _, v := range result {
		if v.Name == n {
			return v, nil
		}
	}

	return nil, not_found.New(constant.ChoiceSetKind, n)
}
