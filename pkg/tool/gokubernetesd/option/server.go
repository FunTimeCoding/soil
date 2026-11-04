package option

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"

type Server struct {
	Address         string
	ServiceTokens   []string
	Arm             bool
	PostgresLocator string
	LitePath        string
	Clusters        func() map[string]*cluster.Cluster
}
