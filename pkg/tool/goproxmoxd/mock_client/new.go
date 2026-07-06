package mock_client

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/luthermonson/go-proxmox"
)

func New() *Client {
	return &Client{
		nodes:              make(map[string]*proxmox.Node),
		nodeStatuses:       make(map[string]*node_status.Status),
		machines:           make(map[string]map[int]*proxmox.VirtualMachine),
		containers:         make(map[string]map[int]*proxmox.Container),
		machineSnapshots:   make(map[int][]*proxmox.VirtualMachineSnapshot),
		containerSnapshots: make(map[int][]*proxmox.ContainerSnapshot),
		networks:           make(map[string]proxmox.NodeNetworks),
		storages:           make(map[string]proxmox.Storages),
		storageContent:     make(map[string][]*proxmox.StorageContent),
		nextID:             100,
	}
}
