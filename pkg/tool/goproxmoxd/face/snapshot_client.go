package face

import "github.com/luthermonson/go-proxmox"

type SnapshotClient interface {
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
}
