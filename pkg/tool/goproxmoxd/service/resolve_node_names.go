package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) resolveNodeNames(
	c face.ProxmoxClient,
	node string,
) ([]string, error) {
	if node != "" {
		return []string{node}, nil
	}

	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	var result []string

	for _, n := range nodes {
		result = append(result, n.Node)
	}

	return result, nil
}
