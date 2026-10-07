package collector

import "github.com/luthermonson/go-proxmox"

func (c *Collector) SetNode(
	hypervisor string,
	r *proxmox.ClusterResource,
) {
	label := []string{hypervisor, r.Node}
	c.node.Status.WithLabelValues(hypervisor, r.Node, r.Status).Set(1)
	c.node.Processor.WithLabelValues(label...).Set(r.CPU)
	c.node.ProcessorCount.WithLabelValues(label...).Set(float64(r.MaxCPU))
	c.node.MemoryUsed.WithLabelValues(label...).Set(float64(r.Mem))
	c.node.MemoryTotal.WithLabelValues(label...).Set(float64(r.MaxMem))
	c.node.DiskUsed.WithLabelValues(label...).Set(float64(r.Disk))
	c.node.DiskTotal.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.node.Uptime.WithLabelValues(label...).Set(float64(r.Uptime))
}
