package face

import "github.com/luthermonson/go-proxmox"

type ClusterClient interface {
	Cluster() (*proxmox.Cluster, error)
	Version() (*proxmox.Version, error)
	MustCluster() *proxmox.Cluster
	MustVersion() *proxmox.Version
}
