package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ResolveCronJobPods(
	x context.Context,
	c *cluster.Cluster,
	name string,
	namespace string,
) ([]string, error) {
	jobs, e := c.Clientset().BatchV1().Jobs(namespace).List(
		x,
		v1.ListOptions{
			LabelSelector: fmt.Sprintf("job-name=%s", name),
		},
	)

	if e != nil {
		return nil, e
	}

	var result []string

	for _, job := range jobs.Items {
		pods, f := ListPodNames(
			x,
			c,
			namespace,
			fmt.Sprintf("job-name=%s", job.Name),
		)

		if f != nil {
			return nil, f
		}

		result = append(result, pods...)
	}

	return result, nil
}
