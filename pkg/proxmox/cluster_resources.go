package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) ClusterResources(
	filters ...string,
) (proxmox.ClusterResources, error) {
	cluster, e := c.Cluster()

	if e != nil {
		return nil, e
	}

	return cluster.Resources(c.context, filters...)
}
