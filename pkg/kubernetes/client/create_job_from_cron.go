package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	batch "k8s.io/api/batch/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateJobFromCron(
	namespace string,
	cron string,
	name string,
) *batch.Job {
	result, e := c.client.BatchV1().Jobs(namespace).Create(
		c.context,
		&batch.Job{
			ObjectMeta: meta.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    map[string]string{"created-by": "manual-trigger"},
			},
			Spec: c.CronJob(namespace, cron).Raw.Spec.JobTemplate.Spec,
		},
		meta.CreateOptions{},
	)
	errors.PanicOnError(e)

	return result
}
