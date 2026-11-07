package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Server struct {
	option.Web
	Arm             bool
	PostgresLocator string
	LitePath        string
	Clusters        func() map[string]*cluster.Cluster
}
