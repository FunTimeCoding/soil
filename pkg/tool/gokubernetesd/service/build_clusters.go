package service

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func buildClusters() map[string]*cluster.Cluster {
	result := make(map[string]*cluster.Cluster)

	if c, e := client.TryInCluster("in-cluster"); e == nil {
		result["in-cluster"] = cluster.New(
			"in-cluster",
			c.Nested(),
			c.Dynamic(),
			c.Configuration(),
		)
	}

	contexts, e := client_configuration.Contexts()

	if e != nil {
		return result
	}

	for _, name := range contexts {
		if _, exists := result[name]; exists {
			continue
		}

		c, f := client.NewContext(name)

		if f != nil {
			continue
		}

		result[name] = cluster.New(
			name,
			c.Nested(),
			c.Dynamic(),
			c.Configuration(),
		)
	}

	return result
}
