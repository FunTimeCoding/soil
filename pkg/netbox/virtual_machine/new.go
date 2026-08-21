package virtual_machine

import (
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
)

func New(v *netbox.VirtualMachineWithConfigContext) *Machine {
	var cluster string

	if v.Cluster.IsSet() {
		cluster = v.Cluster.Get().GetName()
	}

	var site string

	if v.Site.IsSet() {
		site = v.Site.Get().GetName()
	}

	var address string

	if v.PrimaryIp4.IsSet() {
		address = v.PrimaryIp4.Get().GetDisplay()
	}

	status := v.GetStatus()

	return &Machine{
		Identifier:     v.GetId(),
		Name:           v.GetName(),
		Cluster:        cluster,
		Site:           site,
		Status:         string(status.GetValue()),
		PrimaryAddress: address,
		Tags:           tag.Names(v.Tags),
		Raw:            v,
	}
}
