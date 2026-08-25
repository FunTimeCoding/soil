package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *Collector) Clear(hypervisor string) {
	label := prometheus.Labels{constant.HypervisorLabel: hypervisor}

	for _, v := range c.clearable {
		v.DeletePartialMatch(label)
	}
}
