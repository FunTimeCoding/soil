package rack

import (
	"github.com/funtimecoding/soil/pkg/netbox/helper"
	"github.com/netbox-community/go-netbox/v4"
)

func New(v *netbox.Rack) *Rack {
	return &Rack{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Link:       helper.ToWebLink(v.GetUrl()),
		Raw:        v,
	}
}
