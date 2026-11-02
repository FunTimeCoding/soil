package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Nodes() (proxmox.NodeStatuses, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	var result proxmox.NodeStatuses

	for name := range c.nodes {
		result = append(
			result,
			&proxmox.NodeStatus{Node: name, Status: "online"},
		)
	}

	return result, nil
}
