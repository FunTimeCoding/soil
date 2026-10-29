package cluster

func (c *Cluster) Host() string {
	if c.configuration == nil {
		return ""
	}

	return c.configuration.Host
}
