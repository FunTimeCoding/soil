package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
)

func printLog(
	k *client.Client,
	namespace string,
	job string,
) {
	fmt.Println("Fetch job log")
	j := k.Job(namespace, job)

	if j == nil {
		fmt.Printf("Not found: %s/%s\n", namespace, job)

		return
	}

	if len(j.Raw.Status.Conditions) > 0 {
		for _, p := range j.Raw.Status.Conditions {
			if p.Type != "Complete" {
				continue
			}

			if p.Status != "True" {
				continue
			}

			t := j.Raw.Status.CompletionTime.String()
			fmt.Printf("Completed: %s\n", t)
		}
	}

	if j.Raw.Status.Active > 0 {
		fmt.Printf("Active: %d\n", j.Raw.Status.Active)
	}

	if j.Raw.Status.Failed > 0 {
		fmt.Printf("Failed: %d\n", j.Raw.Status.Failed)
	}

	if j.Raw.Status.Succeeded > 0 {
		fmt.Printf("Succeeded: %d\n", j.Raw.Status.Succeeded)
	}

	fmt.Printf("Job name: %s\n", j.Name)
	controller := string(j.Raw.UID)

	for _, p := range k.Pods(
		filter.New().AddNamespaces(namespace).AddNames(j.Name),
	) {
		if p.Label("batch.kubernetes.io/controller-uid") != controller {
			continue
		}

		fmt.Printf("Pod: %s\n", p.Name)
		log := k.Log(namespace, p.Name, "")
		fmt.Println("Log:")
		fmt.Println(log)
		fmt.Println("-----")
	}
}
