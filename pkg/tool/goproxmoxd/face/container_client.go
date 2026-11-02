package face

import "github.com/luthermonson/go-proxmox"

type ContainerClient interface {
	Containers(n *proxmox.Node) (proxmox.Containers, error)
	Container(
		n *proxmox.Node,
		identifier int,
	) (*proxmox.Container, error)
	StartContainer(v *proxmox.Container) (*proxmox.Task, error)
	StopContainer(v *proxmox.Container) (*proxmox.Task, error)
	ShutdownContainer(v *proxmox.Container) (*proxmox.Task, error)
}
