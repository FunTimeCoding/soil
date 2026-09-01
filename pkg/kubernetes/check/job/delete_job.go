package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
)

func deleteJob(
	k *client.Client,
	namespace string,
	name string,
) {
	console.Format("Delete job %s in %s\n", name, namespace)

	if j := k.Job(namespace, name); j != nil {
		console.Format("  %s\n", j.Name)
		k.DeleteJobWatch(namespace, j.Name)
	}
}
