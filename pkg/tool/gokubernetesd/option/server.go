package option

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"

type Server struct {
	Address         string
	ServiceTokens   []string
	ReadOnly        bool
	PostgresLocator string
	LitePath        string
	Clusters        func() map[string]*cluster.Cluster
}
