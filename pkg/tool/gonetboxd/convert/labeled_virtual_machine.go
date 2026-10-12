package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func LabeledVirtualMachine(
	m *virtual_machine.Machine,
	labels []*label.Label,
) *server.VirtualMachine {
	result := VirtualMachine(m)
	result.Labels = new(Labels(labels))

	return result
}
