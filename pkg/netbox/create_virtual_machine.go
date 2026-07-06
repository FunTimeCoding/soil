package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualMachine(
	name string,
	cl *cluster.Cluster,
) (*virtual_machine.Machine, error) {
	q := netbox.NewWritableVirtualMachineWithConfigContextRequest(name)
	q.SetCluster(
		netbox.BriefClusterRequestAsDeviceWithConfigContextRequestCluster(
			netbox.NewBriefClusterRequest(cl.Name),
		),
	)
	q.SetStatus(netbox.PATCHEDWRITABLEVIRTUALMACHINEWITHCONFIGCONTEXTREQUESTSTATUS_ACTIVE)
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesCreate(
		c.context,
	).WritableVirtualMachineWithConfigContextRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_machine.New(result), nil
}
