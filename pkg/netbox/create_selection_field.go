package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_record"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_request"
	"net/http"
)

func (c *Client) CreateSelectionField(
	name string,
	label string,
	choiceSet string,
) (*custom_field.Field, error) {
	s, e := c.ChoiceSetByName(choiceSet)

	if e != nil {
		return nil, e
	}

	q := custom_field_request.New()
	q.Name = name
	q.Label = label
	q.Kind = constant.MultipleSelection
	q.ObjectTypes = []string{
		constant.DeviceObject,
		constant.VirtualMachineObject,
	}
	q.ChoiceSet = s.Identifier
	result := custom_field_record.New()

	if f := c.rawRequest(
		http.MethodPost,
		constant.CustomFieldPath,
		q,
		result,
	); f != nil {
		return nil, f
	}

	return custom_field.New(result), nil
}
