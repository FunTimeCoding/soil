package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ClusterTypes(v []*cluster_type.Type) []*server.ClusterType {
	result := make([]*server.ClusterType, 0, len(v))

	for _, t := range v {
		result = append(result, ClusterType(t))
	}

	return result
}
