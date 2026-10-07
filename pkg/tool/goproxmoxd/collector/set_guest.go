package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector/metric"
	"github.com/luthermonson/go-proxmox"
	"strconv"
	"strings"
)

func (c *Collector) SetGuest(
	hypervisor string,
	r *proxmox.ClusterResource,
) {
	label := []string{
		hypervisor,
		r.Node,
		r.Type,
		strconv.FormatUint(r.VMID, 10),
		r.Name,
	}
	c.guest.Status.WithLabelValues(
		metric.WithLabel(label, r.Status)...,
	).Set(1)
	c.guest.Template.WithLabelValues(label...).Set(float64(r.Template))
	c.guest.Processor.WithLabelValues(label...).Set(r.CPU)
	c.guest.ProcessorCount.WithLabelValues(label...).Set(float64(r.MaxCPU))
	c.guest.MemoryUsed.WithLabelValues(label...).Set(float64(r.Mem))
	c.guest.MemoryTotal.WithLabelValues(label...).Set(float64(r.MaxMem))
	c.guest.DiskUsed.WithLabelValues(label...).Set(float64(r.Disk))
	c.guest.DiskTotal.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.guest.Uptime.WithLabelValues(label...).Set(float64(r.Uptime))
	c.guest.NetworkReceive.WithLabelValues(label...).Set(float64(r.NetIn))
	c.guest.NetworkTransmit.WithLabelValues(label...).Set(float64(r.NetOut))
	c.guest.DiskRead.WithLabelValues(label...).Set(float64(r.DiskRead))
	c.guest.DiskWritten.WithLabelValues(label...).Set(float64(r.DiskWrite))

	for _, tag := range strings.Split(r.Tags, ",") {
		if tag == "" {
			continue
		}

		c.guest.Tag.WithLabelValues(
			metric.WithLabel(label, tag)...,
		).Set(1)
	}
}
