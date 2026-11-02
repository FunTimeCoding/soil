package face

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"
	"github.com/luthermonson/go-proxmox"
)

type MachineService interface {
	ListMachines(
		instance string,
		node string,
	) (proxmox.VirtualMachines, error)
	GetMachine(
		instance string,
		identifier int,
		node string,
	) (*proxmox.VirtualMachine, error)
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
}
