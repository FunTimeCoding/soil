package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestVirtualizationEntities(t *testing.T) {
	assert.NotNil(t, cluster_group.New(&netbox.ClusterGroup{}))
	assert.NotNil(t, cluster.New(&netbox.Cluster{}))
	assert.NotNil(t, cluster_type.New(&netbox.ClusterType{}))
	assert.NotNil(
		t,
		virtual_machine.New(&netbox.VirtualMachineWithConfigContext{}),
	)
}
