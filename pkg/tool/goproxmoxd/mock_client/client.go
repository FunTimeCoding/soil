package mock_client

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/luthermonson/go-proxmox"
	"sync"
)

type Client struct {
	nodes              map[string]*proxmox.Node
	nodeStatuses       map[string]*node_status.Status
	machines           map[string]map[int]*proxmox.VirtualMachine
	containers         map[string]map[int]*proxmox.Container
	machineSnapshots   map[int][]*proxmox.VirtualMachineSnapshot
	containerSnapshots map[int][]*proxmox.ContainerSnapshot
	networks           map[string]proxmox.NodeNetworks
	storages           map[string]proxmox.Storages
	storageContent     map[string][]*proxmox.StorageContent
	nextID             int
	mutex              sync.Mutex
}
