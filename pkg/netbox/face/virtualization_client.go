package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type VirtualizationClient interface {
	ClusterByName(n string) (*cluster.Cluster, error)
	Clusters() ([]*cluster.Cluster, error)
	ClusterTypeByName(n string) (*cluster_type.Type, error)
	ClusterTypes() ([]*cluster_type.Type, error)
	CreateCluster(
		name string,
		t *cluster_type.Type,
		s *site.Site,
	) (*cluster.Cluster, error)
	CreateClusterType(name string) (*cluster_type.Type, error)
	UpdateVirtualInterface(
		vm *virtual_machine.Machine,
		name string,
		h net.HardwareAddr,
	) (*netbox.VMInterface, error)
	CreateVirtualAddress(
		interfaceIdentifier int32,
		address string,
		status string,
	) (*internet_address.Address, error)
	CreateVirtualDisk(
		machine string,
		name string,
		size int32,
	) (*virtual_disk.Disk, error)
	CreateVirtualInterface(
		vm *virtual_machine.Machine,
		name string,
	) (*netbox.VMInterface, error)
	CreateVirtualMachine(
		name string,
		cl *cluster.Cluster,
	) (*virtual_machine.Machine, error)
	RenameVirtualMachine(
		name string,
		newName string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineCores(
		name string,
		cores float64,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineMemory(
		name string,
		megabytes int32,
	) (*virtual_machine.Machine, error)
	SetVirtualMachinePlatform(
		name string,
		platformName string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachinePrimaryAddress(
		name string,
		address string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineRole(
		name string,
		roleName string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineStatus(
		name string,
		status string,
	) (*virtual_machine.Machine, error)
	SetVirtualMachineTenant(
		name string,
		tenantName string,
	) (*virtual_machine.Machine, error)
	VirtualMachineAddresses(machine string) ([]*internet_address.Address, error)
	VirtualMachineByName(n string) (*virtual_machine.Machine, error)
	VirtualMachineDisks(m *virtual_machine.Machine) ([]*virtual_disk.Disk, error)
	VirtualMachineInterfaceByName(
		vm *virtual_machine.Machine,
		name string,
	) (*netbox.VMInterface, error)
	VirtualMachineInterfaces(m *virtual_machine.Machine) ([]*network.Interface, error)
	VirtualMachines() ([]*virtual_machine.Machine, error)
	AssignPhysicalToVirtualInterface(
		p *physical_address.Address,
		identifier int32,
	) (*physical_address.Address, error)
	ClusterGroups() ([]*cluster_group.Group, error)
	ClustersByName(s string) ([]*cluster.Cluster, error)
	DevicesByCluster(s string) ([]*device.Device, error)
	MustClusterGroups() []*cluster_group.Group
	MustClusterTypes() []*cluster_type.Type
	MustClusters() []*cluster.Cluster
	MustVirtualMachines() []*virtual_machine.Machine
}
