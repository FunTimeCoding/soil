package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func LabeledVirtualMachines(
	v []*virtual_machine.Machine,
	grouped map[int32][]*label.Label,
) []*server.VirtualMachine {
	result := make([]*server.VirtualMachine, 0, len(v))

	for _, m := range v {
		entry := VirtualMachine(m)

		if labels, okay := grouped[m.Identifier]; okay {
			entry.Labels = new(Labels(labels))
		}

		result = append(result, entry)
	}

	return result
}
