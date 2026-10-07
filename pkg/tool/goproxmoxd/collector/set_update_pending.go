package collector

func (c *Collector) SetUpdatePending(
	hypervisor string,
	node string,
	count int,
) {
	c.node.UpdatePending.WithLabelValues(hypervisor, node).Set(float64(count))
}
