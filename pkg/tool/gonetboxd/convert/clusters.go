package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Clusters(v []*cluster.Cluster) []*server.Cluster {
	result := make([]*server.Cluster, 0, len(v))

	for _, c := range v {
		result = append(result, Cluster(c))
	}

	return result
}
