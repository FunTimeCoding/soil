package face

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"
	"github.com/luthermonson/go-proxmox"
)

type Service interface {
	Instances() []inventory.Instance
	Instance(name string) (*inventory.Instance, bool)
	ResolveInstance(explicit string) (string, error)
	ActiveInstance(sessionIdentifier string) (string, bool)
	SetActiveInstance(
		sessionIdentifier string,
		instance string,
	)
	Client(instance string) (ProxmoxClient, error)
	SSHClient(instance string) (SnippetClient, error)
	ListNodes(instance string) (proxmox.NodeStatuses, error)
	GetNodeStatus(
		instance string,
		node string,
	) (*node_status.Status, error)
	ListMachines(
		instance string,
		node string,
	) (proxmox.VirtualMachines, error)
	ListContainers(
		instance string,
		node string,
	) (proxmox.Containers, error)
	GetMachine(
		instance string,
		identifier int,
		node string,
	) (*proxmox.VirtualMachine, error)
	GetContainer(
		instance string,
		identifier int,
		node string,
	) (*proxmox.Container, error)
	ListNetworks(
		instance string,
		node string,
	) (proxmox.NodeNetworks, error)
	ListStorages(
		instance string,
		node string,
	) (proxmox.Storages, error)
	ListStorageContent(
		instance string,
		node string,
		storage string,
	) ([]*proxmox.StorageContent, error)
	DownloadLocator(
		instance string,
		node string,
		storage string,
		content string,
		filename string,
		l string,
	) error
	CreateMachine(
		instance string,
		m *create_machine.Machine,
	) (int, error)
	DeriveHardwareAddress(
		instance string,
		identifier int,
	) (string, *int, error)
	UpdateMachine(
		instance string,
		a *update_machine.Machine,
	) error
	CloneMachine(
		instance string,
		identifier int,
		node string,
		options *proxmox.VirtualMachineCloneOptions,
	) (int, error)
	DeleteMachine(
		instance string,
		identifier int,
		node string,
		purge bool,
	) error
	StartMachine(
		instance string,
		identifier int,
		node string,
	) (string, error)
	StopMachine(
		instance string,
		identifier int,
		node string,
	) (string, error)
	ShutdownMachine(
		instance string,
		identifier int,
		node string,
	) (string, error)
	ResetMachine(
		instance string,
		identifier int,
		node string,
	) (string, error)
	StartContainer(
		instance string,
		identifier int,
		node string,
	) (string, error)
	StopContainer(
		instance string,
		identifier int,
		node string,
	) (string, error)
	ShutdownContainer(
		instance string,
		identifier int,
		node string,
	) (string, error)
	ListMachineSnapshots(
		instance string,
		identifier int,
		node string,
	) ([]*proxmox.VirtualMachineSnapshot, error)
	CreateMachineSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
	RollbackMachineSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
	DeleteMachineSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
	ListContainerSnapshots(
		instance string,
		identifier int,
		node string,
	) ([]*proxmox.ContainerSnapshot, error)
	CreateContainerSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
	RollbackContainerSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
	DeleteContainerSnapshot(
		instance string,
		identifier int,
		node string,
		name string,
	) (string, error)
}
