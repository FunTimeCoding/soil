package collector

import "github.com/luthermonson/go-proxmox"

func (c *Collector) SetNode(
	hypervisor string,
	r *proxmox.ClusterResource,
) {
	label := []string{hypervisor, r.Node}
	c.node.status.WithLabelValues(hypervisor, r.Node, r.Status).Set(1)
	c.node.processor.WithLabelValues(label...).Set(r.CPU)
	c.node.processorCount.WithLabelValues(label...).Set(float64(r.MaxCPU))
	c.node.memoryUsed.WithLabelValues(label...).Set(float64(r.Mem))
	c.node.memoryTotal.WithLabelValues(label...).Set(float64(r.MaxMem))
	c.node.diskUsed.WithLabelValues(label...).Set(float64(r.Disk))
	c.node.diskTotal.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.node.uptime.WithLabelValues(label...).Set(float64(r.Uptime))
}
