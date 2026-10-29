package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_record"
)

func (c *Client) CreateConfigurationContext(
	name string,
	weight int32,
	payload map[string]any,
	a *configuration_assignment.Assignment,
) (*configuration_context.Context, error) {
	v := configuration_record.New()
	v.Identifier = c.nextIdentifier()
	v.Name = name
	v.Weight = weight
	v.Active = true
	v.Payload = payload
	result := configuration_context.New(v)
	assign(result, a)
	c.contexts = append(c.contexts, result)
	c.record("CreateConfigurationContext", name)

	return result, nil
}
