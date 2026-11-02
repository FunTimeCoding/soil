package service

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
)

func (s *Service) sampleClient(
	c face.ProxmoxClient,
	result *sample.Sample,
) error {
	resources, e := c.ClusterResources()

	if e != nil {
		return e
	}

	result.Resources = resources

	for _, r := range resources {
		if r.Type == proxmoxConstant.NodeType {
			result.Nodes = append(
				result.Nodes,
				s.sampleNode(c, result.Hypervisor, r.Node),
			)
		}
	}

	result.Unbacked, e = c.GuestsNotInBackup()

	return e
}
