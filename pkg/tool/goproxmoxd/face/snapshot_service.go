package face

import "github.com/luthermonson/go-proxmox"

type SnapshotService interface {
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
