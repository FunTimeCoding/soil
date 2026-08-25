package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func guestResource(
	guestType string,
	node string,
	identifier int,
	name string,
	status string,
	processorCount int,
	memory uint64,
	maximumMemory uint64,
) *proxmox.ClusterResource {
	return &proxmox.ClusterResource{
		ID:     fmt.Sprintf("%s/%d", guestType, identifier),
		Type:   guestType,
		Node:   node,
		VMID:   uint64(identifier),
		Name:   name,
		Status: status,
		MaxCPU: uint64(processorCount),
		Mem:    memory,
		MaxMem: maximumMemory,
	}
}
