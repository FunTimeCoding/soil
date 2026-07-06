package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func VirtualMachine(m *virtual_machine.Machine) *server.VirtualMachine {
	result := &server.VirtualMachine{Identifier: m.Identifier, Name: m.Name}

	if m.Raw.Cluster.IsSet() && m.Raw.Cluster.Get() != nil {
		result.Cluster = new(m.Raw.Cluster.Get().GetName())
	}

	if m.Raw.Site.IsSet() && m.Raw.Site.Get() != nil {
		result.Site = new(m.Raw.Site.Get().GetName())
	}

	if len(m.Tags) > 0 {
		result.Tags = &m.Tags
	}

	return result
}
