package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (c *Client) ChoiceSets() ([]*choice_set.Set, error) {
	return c.sets, nil
}
