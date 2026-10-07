package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector/metric"
	"github.com/luthermonson/go-proxmox"
)

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
	c.storage.Status.WithLabelValues(
		metric.WithLabel(label, r.Status)...,
	).Set(1)
	c.storage.Used.WithLabelValues(label...).Set(float64(r.Disk))
	c.storage.Total.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.storage.Shared.WithLabelValues(label...).Set(float64(r.Shared))
}
