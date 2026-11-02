package face

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/luthermonson/go-proxmox"
)

type NodeClient interface {
	Nodes() (proxmox.NodeStatuses, error)
	Node(name string) (*proxmox.Node, error)
	NodeStatus(name string) (*node_status.Status, error)
	NodeVersion(n *proxmox.Node) (*proxmox.Version, error)
	ClusterResources(filters ...string) (proxmox.ClusterResources, error)
	GuestsNotInBackup() ([]*proxmox.BackupGuestEntry, error)
	UpdatesPending(n *proxmox.Node) ([]*proxmox.APTUpdate, error)
	Networks(n *proxmox.Node) (proxmox.NodeNetworks, error)
	Storages(n *proxmox.Node) (proxmox.Storages, error)
	Storage(
		n *proxmox.Node,
		name string,
	) (*proxmox.Storage, error)
	StorageContent(s *proxmox.Storage) ([]*proxmox.StorageContent, error)
	DownloadLocator(
		s *proxmox.Storage,
		content string,
		filename string,
		l string,
	) (*proxmox.Task, error)
}
