package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"k8s.io/api/batch/v1"
)

func createJob(
	k *client.Client,
	namespace string,
	cron string,
	name string,
) *v1.Job {
	console.Format("Create job %s in %s from %s\n", name, namespace, cron)
	result := k.CreateJobFromCron(namespace, cron, name)
	console.Format("  %s\n", result.Name)

	return result
}
