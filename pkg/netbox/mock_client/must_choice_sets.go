package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
)

func (c *Client) MustChoiceSets() []*choice_set.Set {
	result, e := c.ChoiceSets()
	errors.PanicOnError(e)

	return result
}
