package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_record"
)

func (c *Client) CreateChoiceSet(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	v := choice_set_record.New()
	v.Identifier = c.nextIdentifier()
	v.Name = name
	result := choice_set.New(v)
	result.Choices = choices
	c.sets = append(c.sets, result)
	c.record("CreateChoiceSet", name)

	return result, nil
}
