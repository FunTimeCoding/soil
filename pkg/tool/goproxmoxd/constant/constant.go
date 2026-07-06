package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
)

var ErrorCDROMCloudInitConflict = errors.New(
	"cdrom and cloud-init are mutually exclusive - both use ide2",
)

var (
	ErrorNoChanges         = errors.New("no changes specified")
	ErrorMachineRunning    = errors.New("vm is running - stop it before deleting")
	ErrorSetDeleteConflict = errors.New("cannot set and delete the same field")
)
var Identity = identity.New(
	"goproxmoxd",
	"Proxmox hypervisor bridge for node and VM inspection",
	"goproxmoxd",
).WithInstructions(
	"Proxmox hypervisor bridge.",
	paragraph.New(
		MultiInstance,
		"Multi-instance - call list_instances and use_instance before querying.",
	),
)

const MultiInstance = "multi_instance"
const SnippetDirectory = "/var/lib/vz/snippets"
const (
	ListInstances             = "list_instances"
	UseInstance               = "use_instance"
	ListNodes                 = "list_nodes"
	GetNodeStatus             = "get_node_status"
	ListMachines              = "list_machines"
	GetMachine                = "get_machine"
	StartMachine              = "start_machine"
	StopMachine               = "stop_machine"
	ShutdownMachine           = "shutdown_machine"
	ResetMachine              = "reset_machine"
	ListMachineSnapshots      = "list_machine_snapshots"
	CreateMachineSnapshot     = "create_machine_snapshot"
	RollbackMachineSnapshot   = "rollback_machine_snapshot"
	DeleteMachineSnapshot     = "delete_machine_snapshot"
	ListContainers            = "list_containers"
	GetContainer              = "get_container"
	ListContainerSnapshots    = "list_container_snapshots"
	CreateContainerSnapshot   = "create_container_snapshot"
	RollbackContainerSnapshot = "rollback_container_snapshot"
	DeleteContainerSnapshot   = "delete_container_snapshot"
	ListNetworks              = "list_networks"
	ListStorages              = "list_storages"
	ListStorageContent        = "list_storage_content"
	CreateSnippet             = "create_snippet"
	GetSnippet                = "get_snippet"
	ListSnippets              = "list_snippets"
	DeleteSnippet             = "delete_snippet"
	DownloadLocator           = "download_locator"
	CreateMachine             = "create_machine"
	UpdateMachine             = "update_machine"
	CloneMachine              = "clone_machine"
	DeleteMachine             = "delete_machine"
)
