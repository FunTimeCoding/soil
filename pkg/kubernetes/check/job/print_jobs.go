package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
)

func printJobs(
	k *client.Client,
	namespace string,
) {
	fmt.Printf("Jobs in %s:\n", namespace)

	for _, j := range k.Jobs(namespace) {
		fmt.Printf("  %s\n", j.Name)
	}
}
