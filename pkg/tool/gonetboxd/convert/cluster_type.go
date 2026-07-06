package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ClusterType(t *cluster_type.Type) *server.ClusterType {
	return &server.ClusterType{Identifier: t.Identifier, Name: t.Name}
}
