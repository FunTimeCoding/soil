package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
)

func printJobs(
	k *client.Client,
	namespace string,
) {
	console.Format("Jobs in %s:\n", namespace)

	for _, j := range k.Jobs(namespace) {
		console.Format("  %s\n", j.Name)
	}
}
