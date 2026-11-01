package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"time"
)

const (
	PollInterval      = 30 * time.Second
	ReconcileInterval = time.Hour
)
const (
	FloorTitle = "Floor"
	FloorPath  = "/"

	FloorEvent   = "floor"
	SummaryEvent = "summary"

	StorageWarnRatio = 0.8

	ReleaseClass    = "release"
	NodeHeaderClass = "node-header"
	FloorTableClass = "floor-table"
	StoppedClass    = "stopped"
	LoadClass       = "load"
)

var Identity = identity.New(
	"goproxmoxd",
	"Proxmox hypervisor bridge for node and VM inspection",
	"goproxmoxd",
).WithInstructions(
	"Proxmox hypervisor bridge.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const SnippetDirectory = "/var/lib/vz/snippets"
const (
	ListNodes                 = "list_nodes"
	GetNodeStatus             = "get_node_status"
	ListMachines              = "list_machines"
	GetMachine                = "get_machine"
	DeriveHardwareAddress     = "derive_hardware_address"
	StartMachine              = "start_machine"
	StopMachine               = "stop_machine"
	ShutdownMachine           = "shutdown_machine"
	ResetMachine              = "reset_machine"
	Reconcile                 = "reconcile"
	ListMachineSnapshots      = "list_machine_snapshots"
	CreateMachineSnapshot     = "create_machine_snapshot"
	RollbackMachineSnapshot   = "rollback_machine_snapshot"
	DeleteMachineSnapshot     = "delete_machine_snapshot"
	ListContainers            = "list_containers"
	GetContainer              = "get_container"
	StartContainer            = "start_container"
	StopContainer             = "stop_container"
	ShutdownContainer         = "shutdown_container"
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
const (
	HostEnvironment     = "GOPROXMOX_HOST"
	TokenEnvironment    = "GOPROXMOX_TOKEN"
	PortEnvironment     = "GOPROXMOX_PORT"
	InsecureEnvironment = "GOPROXMOX_INSECURE"
)
