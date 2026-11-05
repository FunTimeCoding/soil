package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (c *Client) ChoiceSets() ([]*choice_set.Set, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.sets, nil
}
