package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"strings"
)

func ResolvePods(
	x context.Context,
	c *cluster.Cluster,
	name string,
	namespace string,
) ([]string, error) {
	parts := strings.SplitN(name, "/", 2)

	if len(parts) == 1 {
		return []string{name}, nil
	}

	resourceType := parts[0]
	resourceName := parts[1]

	switch resourceType {
	case "deployment":
		return ResolveDeploymentPods(x, c, resourceName, namespace)
	case "job":
		return ResolveJobPods(x, c, resourceName, namespace)
	case "cronjob":
		return ResolveCronJobPods(x, c, resourceName, namespace)
	default:
		return nil, fmt.Errorf(
			"unsupported resource type for logs: %s (use deployment, job, or cronjob)",
			resourceType,
		)
	}
}
