package node_sample

import "github.com/luthermonson/go-proxmox"

type NodeSample struct {
	Name           string
	Version        *proxmox.Version
	Updates        []*proxmox.APTUpdate
	UpdatesFailure error
}
