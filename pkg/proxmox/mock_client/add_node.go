package mock_client

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) AddNode(name string) {
	c.nodes[name] = &proxmox.Node{Name: name}
	c.nodeStatuses[name] = node_status.Stub()
	c.machines[name] = make(map[int]*proxmox.VirtualMachine)
	c.containers[name] = make(map[int]*proxmox.Container)
}
