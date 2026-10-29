package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	batch "k8s.io/api/batch/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateJobFromCron(
	namespace string,
	cron string,
	name string,
) (*batch.Job, error) {
	j, e := c.CronJob(namespace, cron)

	if e != nil {
		return nil, e
	}

	result, f := c.client.BatchV1().Jobs(namespace).Create(
		c.context,
		&batch.Job{
			ObjectMeta: meta.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    map[string]string{"created-by": "manual-trigger"},
			},
			Spec: j.Raw.Spec.JobTemplate.Spec,
		},
		meta.CreateOptions{},
	)

	if f != nil {
		return nil, refusal.Classify(f)
	}

	return result, nil
}
