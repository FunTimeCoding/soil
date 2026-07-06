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
	ListNodes(c ProxmoxClient) (proxmox.NodeStatuses, error)
	GetNodeStatus(
		c ProxmoxClient,
		node string,
	) (*node_status.Status, error)
	ListMachines(
		c ProxmoxClient,
		node string,
	) (proxmox.VirtualMachines, error)
	ListContainers(
		c ProxmoxClient,
		node string,
	) (proxmox.Containers, error)
	GetMachine(
		c ProxmoxClient,
		identifier int,
		node string,
	) (*proxmox.VirtualMachine, error)
	GetContainer(
		c ProxmoxClient,
		identifier int,
		node string,
	) (*proxmox.Container, error)
	ListNetworks(
		c ProxmoxClient,
		node string,
	) (proxmox.NodeNetworks, error)
	ListStorages(
		c ProxmoxClient,
		node string,
	) (proxmox.Storages, error)
	ListStorageContent(
		c ProxmoxClient,
		node string,
		storage string,
	) ([]*proxmox.StorageContent, error)
	DownloadLocator(
		c ProxmoxClient,
		node string,
		storage string,
		content string,
		filename string,
		l string,
	) error
	CreateMachine(
		c ProxmoxClient,
		m *create_machine.Machine,
	) (int, error)
	UpdateMachine(
		c ProxmoxClient,
		a *update_machine.Machine,
	) error
	CloneMachine(
		c ProxmoxClient,
		identifier int,
		node string,
		options *proxmox.VirtualMachineCloneOptions,
	) (int, error)
	DeleteMachine(
		c ProxmoxClient,
		identifier int,
		node string,
		purge bool,
	) error
	StartMachine(
		c ProxmoxClient,
		identifier int,
		node string,
	) (string, error)
	StopMachine(
		c ProxmoxClient,
		identifier int,
		node string,
	) (string, error)
	ShutdownMachine(
		c ProxmoxClient,
		identifier int,
		node string,
	) (string, error)
	ResetMachine(
		c ProxmoxClient,
		identifier int,
		node string,
	) (string, error)
	ListMachineSnapshots(
		c ProxmoxClient,
		identifier int,
		node string,
	) ([]*proxmox.VirtualMachineSnapshot, error)
	CreateMachineSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
	RollbackMachineSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
	DeleteMachineSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
	ListContainerSnapshots(
		c ProxmoxClient,
		identifier int,
		node string,
	) ([]*proxmox.ContainerSnapshot, error)
	CreateContainerSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
	RollbackContainerSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
	DeleteContainerSnapshot(
		c ProxmoxClient,
		identifier int,
		node string,
		name string,
	) (string, error)
}
