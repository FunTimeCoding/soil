package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"sort"
)

func (s *Service) Clusters() []*cluster.Cluster {
	var result []*cluster.Cluster

	for _, c := range s.clusters {
		result = append(result, c)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name() < result[j].Name()
		},
	)

	return result
}
