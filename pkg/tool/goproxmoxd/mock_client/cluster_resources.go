package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
	"slices"
)

func (c *Client) ClusterResources(
	filters ...string,
) (proxmox.ClusterResources, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	var result proxmox.ClusterResources

	for name := range c.nodes {
		result = append(
			result,
			&proxmox.ClusterResource{
				ID:     fmt.Sprintf("%s/%s", constant.NodeType, name),
				Type:   constant.NodeType,
				Node:   name,
				Status: constant.OnlineStatus,
			},
		)
	}

	for node, machines := range c.machines {
		for identifier, v := range machines {
			result = append(
				result,
				guestResource(
					constant.MachineType,
					node,
					identifier,
					v.Name,
					v.Status,
					v.CPUs,
					v.Mem,
					v.MaxMem,
				),
			)
		}
	}

	for node, containers := range c.containers {
		for identifier, v := range containers {
			result = append(
				result,
				guestResource(
					constant.ContainerType,
					node,
					identifier,
					v.Name,
					v.Status,
					v.CPUs,
					0,
					v.MaxMem,
				),
			)
		}
	}

	for node, storages := range c.storages {
		for _, s := range storages {
			result = append(
				result,
				&proxmox.ClusterResource{
					ID: fmt.Sprintf(
						"%s/%s/%s",
						constant.StorageType,
						node,
						s.Name,
					),
					Type:       constant.StorageType,
					Node:       node,
					Storage:    s.Name,
					PluginType: s.Type,
					Content:    s.Content,
					Status:     constant.AvailableStatus,
				},
			)
		}
	}

	if len(filters) == 0 {
		return result, nil
	}

	var filtered proxmox.ClusterResources

	for _, r := range result {
		if slices.Contains(filters, r.Type) {
			filtered = append(filtered, r)
		}
	}

	return filtered, nil
}
