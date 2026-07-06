package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPodNames(
	x context.Context,
	c *cluster.Cluster,
	namespace string,
	labelSelector string,
) ([]string, error) {
	pods, e := c.Clientset().CoreV1().Pods(namespace).List(
		x,
		v1.ListOptions{LabelSelector: labelSelector},
	)

	if e != nil {
		return nil, e
	}

	var result []string

	for _, pod := range pods.Items {
		result = append(result, pod.Name)
	}

	return result, nil
}
