package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/intent"
	"github.com/luthermonson/go-proxmox"
)

type NodeService interface {
	ListNodes(instance string) (proxmox.NodeStatuses, error)
	GetNodeStatus(
		instance string,
		node string,
	) (*node_status.Status, error)
	ListNetworks(
		instance string,
		node string,
	) (proxmox.NodeNetworks, error)
	ListStorages(
		instance string,
		node string,
	) (proxmox.Storages, error)
	ListStorageContent(
		instance string,
		node string,
		storage string,
	) ([]*proxmox.StorageContent, error)
	DownloadLocator(
		instance string,
		node string,
		storage string,
		content string,
		filename string,
		l string,
	) error
	Reconcile(
		x context.Context,
		hypervisor string,
		full bool,
		write bool,
	) (*intent.Plan, error)
}
