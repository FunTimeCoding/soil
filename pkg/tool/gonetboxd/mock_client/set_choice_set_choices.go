package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (c *Client) SetChoiceSetChoices(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	v, e := c.ChoiceSetByName(name)

	if e != nil {
		return nil, e
	}

	v.Choices = choices
	c.record("SetChoiceSetChoices", name)

	return v, nil
}
