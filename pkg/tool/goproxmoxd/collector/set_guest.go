package collector

import (
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
	c.guest.status.WithLabelValues(
		withLabel(label, r.Status)...,
	).Set(1)
	c.guest.template.WithLabelValues(label...).Set(float64(r.Template))
	c.guest.processor.WithLabelValues(label...).Set(r.CPU)
	c.guest.processorCount.WithLabelValues(label...).Set(float64(r.MaxCPU))
	c.guest.memoryUsed.WithLabelValues(label...).Set(float64(r.Mem))
	c.guest.memoryTotal.WithLabelValues(label...).Set(float64(r.MaxMem))
	c.guest.diskUsed.WithLabelValues(label...).Set(float64(r.Disk))
	c.guest.diskTotal.WithLabelValues(label...).Set(float64(r.MaxDisk))
	c.guest.uptime.WithLabelValues(label...).Set(float64(r.Uptime))
	c.guest.networkReceive.WithLabelValues(label...).Set(float64(r.NetIn))
	c.guest.networkTransmit.WithLabelValues(label...).Set(float64(r.NetOut))
	c.guest.diskRead.WithLabelValues(label...).Set(float64(r.DiskRead))
	c.guest.diskWritten.WithLabelValues(label...).Set(float64(r.DiskWrite))

	for _, tag := range strings.Split(r.Tags, ",") {
		if tag == "" {
			continue
		}

		c.guest.tag.WithLabelValues(
			withLabel(label, tag)...,
		).Set(1)
	}
}
