package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
)

func deleteJob(
	k *client.Client,
	namespace string,
	name string,
) {
	fmt.Printf("Delete job %s in %s\n", name, namespace)

	if j := k.Job(namespace, name); j != nil {
		fmt.Printf("  %s\n", j.Name)
		k.DeleteJobWatch(namespace, j.Name)
	}
}
