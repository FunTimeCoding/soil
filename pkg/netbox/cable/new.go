package cable

import (
	"github.com/funtimecoding/soil/pkg/netbox/helper"
	"github.com/netbox-community/go-netbox/v4"
)

func New(v *netbox.Cable) *Cable {
	var status string

	if v.Status != nil {
		status = string(v.Status.GetValue())
	}

	return &Cable{
		Identifier:  v.GetId(),
		Name:        v.GetDisplay(),
		Description: v.GetDescription(),
		Status:      status,
		SideA:       terminationLabel(v.GetATerminations()),
		SideB:       terminationLabel(v.GetBTerminations()),
		Link:        helper.ToWebLink(v.GetUrl()),
		Raw:         v,
	}
}
