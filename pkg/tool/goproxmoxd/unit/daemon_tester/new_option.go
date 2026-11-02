package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"maps"
	"slices"
)

func newOption(clients map[string]*mock_client.Client) *option.Proxmox {
	var instances []instance.Instance

	for index, name := range slices.Sorted(maps.Keys(clients)) {
		instances = append(
			instances,
			instance.Instance{
				Name:  name,
				Host:  join.Empty(name, ".example"),
				Index: index,
			},
		)
	}

	result := option.New()
	result.Inventory = inventory.New[instance.Instance](instances...)
	result.Hypervisor = func(i *instance.Instance) (proxmox.ProxmoxClient, error) {
		return clients[i.Name], nil
	}
	result.Shell = func(*instance.Instance) (ssh.SSHClient, error) {
		return sshMock.New(), nil
	}

	return result
}
