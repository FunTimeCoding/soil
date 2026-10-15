package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_record"
)

func (c *Client) CreateSelectionField(
	name string,
	label string,
	choiceSet string,
) (*custom_field.Field, error) {
	if _, e := c.ChoiceSetByName(choiceSet); e != nil {
		return nil, e
	}

	v := custom_field_record.New()
	v.Identifier = c.nextIdentifier()
	v.Name = name
	v.Label = label
	v.Kind.Value = constant.MultipleSelection
	v.ObjectTypes = []string{
		constant.DeviceObject,
		constant.VirtualMachineObject,
	}
	result := custom_field.New(v)
	result.ChoiceSet = choiceSet
	c.fields = append(c.fields, result)
	c.record("CreateSelectionField", name, choiceSet)

	return result, nil
}
