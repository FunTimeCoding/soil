package face

import "github.com/luthermonson/go-proxmox"

type ContainerService interface {
	ListContainers(
		instance string,
		node string,
	) (proxmox.Containers, error)
	GetContainer(
		instance string,
		identifier int,
		node string,
	) (*proxmox.Container, error)
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
}
