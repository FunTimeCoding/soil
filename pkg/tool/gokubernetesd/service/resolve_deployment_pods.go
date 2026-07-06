package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ResolveDeploymentPods(
	x context.Context,
	c *cluster.Cluster,
	name string,
	namespace string,
) ([]string, error) {
	deployment, e := c.Clientset().AppsV1().Deployments(namespace).Get(
		x,
		name,
		v1.GetOptions{},
	)

	if e != nil {
		return nil, e
	}

	var selectors []string

	for key, value := range deployment.Spec.Selector.MatchLabels {
		selectors = append(selectors, fmt.Sprintf("%s=%s", key, value))
	}

	return ListPodNames(x, c, namespace, join.Comma(selectors))
}
