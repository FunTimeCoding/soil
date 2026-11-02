package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Nodes() (proxmox.NodeStatuses, error) {
	var result proxmox.NodeStatuses

	for name := range c.nodes {
		result = append(
			result,
			&proxmox.NodeStatus{Node: name, Status: "online"},
		)
	}

	return result, nil
}
