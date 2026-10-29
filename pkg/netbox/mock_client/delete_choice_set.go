package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"slices"
)

func (c *Client) DeleteChoiceSet(name string) error {
	v, e := c.ChoiceSetByName(name)

	if e != nil {
		return e
	}

	c.sets = slices.DeleteFunc(
		c.sets,
		func(x *choice_set.Set) bool { return x == v },
	)
	c.record("DeleteChoiceSet", name)

	return nil
}
