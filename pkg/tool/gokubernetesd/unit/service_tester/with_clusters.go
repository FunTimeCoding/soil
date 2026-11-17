package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"

func WithClusters(clusters map[string]*cluster.Cluster) Option {
	return func(o *Tester) {
		o.clusters = clusters
	}
}
