package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type VirtualizationClient interface {
	AddVirtualJournalEntry(
		x context.Context,
		machine string,
		kind string,
		comments string,
	) (client.JournalEntry, error)
	AddVirtualTag(
		x context.Context,
		name string,
		tag string,
	) (client.VirtualMachine, error)
	CreateCluster(
		x context.Context,
		name string,
		clusterType string,
		site string,
	) (client.Cluster, error)
	CreateClusterType(
		x context.Context,
		name string,
	) (client.ClusterType, error)
	CreateVirtualAddress(
		x context.Context,
		vmName string,
		interfaceName string,
		address string,
		status string,
	) (client.Address, error)
	CreateVirtualDisk(
		x context.Context,
		machine string,
		name string,
		size int32,
	) (client.VirtualDisk, error)
	CreateVirtualInterface(
		x context.Context,
		vmName string,
		name string,
	) (client.VirtualInterface, error)
	CreateVirtualMachine(
		x context.Context,
		name string,
		cluster string,
	) (client.VirtualMachine, error)
	CreateVirtualTunnelTermination(
		x context.Context,
		vm string,
		tunnel string,
		interfaceName string,
		role string,
	) (client.TunnelTermination, error)
	GetVirtualMachine(
		x context.Context,
		name string,
	) (client.VirtualMachine, error)
	ListClusterTypes(
		x context.Context,
	) ([]*client.ClusterType, error)
	ListClusters(
		x context.Context,
	) ([]*client.Cluster, error)
	ListVirtualAddresses(
		x context.Context,
		name string,
	) ([]*client.Address, error)
	ListVirtualDisks(
		x context.Context,
		machine string,
	) ([]*client.VirtualDisk, error)
	ListVirtualInterfaces(
		x context.Context,
		name string,
	) ([]*client.Interface, error)
	ListVirtualJournalEntries(
		x context.Context,
		machine string,
		limit int32,
		offset int32,
	) ([]*client.JournalEntry, error)
	ListVirtualLabels(
		x context.Context,
		machine string,
	) ([]*client.Label, error)
	ListVirtualMachines(
		x context.Context,
	) ([]*client.VirtualMachine, error)
	RemoveVirtualLabel(
		x context.Context,
		machine string,
		key string,
	) error
	RemoveVirtualTag(
		x context.Context,
		name string,
		tag string,
	) (client.VirtualMachine, error)
	SetVirtualInterfacePhysicalAddress(
		x context.Context,
		machine string,
		name string,
		address string,
	) (client.PhysicalAddress, error)
	SetVirtualLabel(
		x context.Context,
		machine string,
		key string,
		value string,
	) (client.Label, error)
	UpdateVirtualMachine(
		x context.Context,
		name string,
		body client.UpdateVirtualMachineRequest,
	) (client.VirtualMachine, error)
}
