package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
)

func (t *Tester) daemonOption() *option.Server {
	result := option.New()
	result.LitePath = t.litePath
	result.Arm = t.arm
	result.Clusters = func() map[string]*cluster.Cluster { return t.clusters }

	return result
}
