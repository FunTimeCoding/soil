package face

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/luthermonson/go-proxmox"
)

type ProxmoxClient interface {
	Nodes() (proxmox.NodeStatuses, error)
	Node(name string) (*proxmox.Node, error)
	NodeStatus(name string) (*node_status.Status, error)
	Machines(n *proxmox.Node) (proxmox.VirtualMachines, error)
	Machine(
		n *proxmox.Node,
		identifier int,
	) (*proxmox.VirtualMachine, error)
	CreateMachine(
		n *proxmox.Node,
		identifier int,
		options ...proxmox.VirtualMachineOption,
	) (*proxmox.Task, error)
	CloneMachine(
		vm *proxmox.VirtualMachine,
		options *proxmox.VirtualMachineCloneOptions,
	) (int, *proxmox.Task, error)
	DeleteMachine(
		v *proxmox.VirtualMachine,
		options *proxmox.VirtualMachineDeleteOptions,
	) (*proxmox.Task, error)
	StartMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error)
	StopMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error)
	ShutdownMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error)
	ResetMachine(v *proxmox.VirtualMachine) (*proxmox.Task, error)
	ResizeDisk(
		v *proxmox.VirtualMachine,
		disk string,
		size string,
	) (*proxmox.Task, error)
	UpdateMachineConfiguration(
		v *proxmox.VirtualMachine,
		options ...proxmox.VirtualMachineOption,
	) error
	NextIdentifier() (int, error)
	WaitForTask(
		t *proxmox.Task,
		seconds int,
	) error
	MachineSnapshots(v *proxmox.VirtualMachine) ([]*proxmox.VirtualMachineSnapshot, error)
	CreateMachineSnapshot(
		v *proxmox.VirtualMachine,
		name string,
	) (*proxmox.Task, error)
	DeleteMachineSnapshot(
		v *proxmox.VirtualMachine,
		name string,
	) (*proxmox.Task, error)
	RollbackMachineSnapshot(
		v *proxmox.VirtualMachine,
		name string,
	) (*proxmox.Task, error)
	Containers(n *proxmox.Node) (proxmox.Containers, error)
	Container(
		n *proxmox.Node,
		identifier int,
	) (*proxmox.Container, error)
	ContainerSnapshots(v *proxmox.Container) ([]*proxmox.ContainerSnapshot, error)
	CreateContainerSnapshot(
		v *proxmox.Container,
		name string,
	) (*proxmox.Task, error)
	DeleteContainerSnapshot(
		v *proxmox.Container,
		name string,
	) (*proxmox.Task, error)
	RollbackContainerSnapshot(
		v *proxmox.Container,
		name string,
	) (*proxmox.Task, error)
	Networks(n *proxmox.Node) (proxmox.NodeNetworks, error)
	Storages(n *proxmox.Node) (proxmox.Storages, error)
	Storage(
		n *proxmox.Node,
		name string,
	) (*proxmox.Storage, error)
	StorageContent(s *proxmox.Storage) ([]*proxmox.StorageContent, error)
	DownloadLocator(
		s *proxmox.Storage,
		content string,
		filename string,
		l string,
	) (*proxmox.Task, error)
}
