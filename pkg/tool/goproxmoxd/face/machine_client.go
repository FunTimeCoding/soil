package face

import "github.com/luthermonson/go-proxmox"

type MachineClient interface {
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
}
