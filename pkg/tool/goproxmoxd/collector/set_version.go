package collector

import "github.com/luthermonson/go-proxmox"

func (c *Collector) SetVersion(
	hypervisor string,
	node string,
	v *proxmox.Version,
) {
	c.node.version.WithLabelValues(
		hypervisor,
		node,
		v.Release,
		v.RepoID,
		v.Version,
	).Set(1)
}
