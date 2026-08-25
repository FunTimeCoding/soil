package collector

import "github.com/luthermonson/go-proxmox"

func (c *Collector) SetStorage(
	hypervisor string,
	r *proxmox.ClusterResource,
) {
	label := []string{
		hypervisor,
		r.Node,
		r.Storage,
		r.PluginType,
		sortedContent(r.Content),
	}
	c.storage.status.WithLabelValues(
		withLabel(label, r.Status)...,
	).Set(1)
	c.storage.used.WithLabelValues(label...).Set(float64(r.Disk))
	c.storage.total.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.storage.shared.WithLabelValues(label...).Set(float64(r.Shared))
}
