package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func ResolveJobPods(
	x context.Context,
	c *cluster.Cluster,
	name string,
	namespace string,
) ([]string, error) {
	return ListPodNames(
		x,
		c,
		namespace,
		fmt.Sprintf("job-name=%s", name),
	)
}
