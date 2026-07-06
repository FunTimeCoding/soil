package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"k8s.io/api/batch/v1"
)

func createJob(
	k *client.Client,
	namespace string,
	cron string,
	name string,
) *v1.Job {
	fmt.Printf("Create job %s in %s from %s\n", name, namespace, cron)
	result := k.CreateJobFromCron(namespace, cron, name)
	fmt.Printf("  %s\n", result.Name)

	return result
}
