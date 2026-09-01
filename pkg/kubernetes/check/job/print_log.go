package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
)

func printLog(
	k *client.Client,
	namespace string,
	job string,
) {
	console.Line("Fetch job log")
	j := k.Job(namespace, job)

	if j == nil {
		console.Format("Not found: %s/%s\n", namespace, job)

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
			console.Format("Completed: %s\n", t)
		}
	}

	if j.Raw.Status.Active > 0 {
		console.Format("Active: %d\n", j.Raw.Status.Active)
	}

	if j.Raw.Status.Failed > 0 {
		console.Format("Failed: %d\n", j.Raw.Status.Failed)
	}

	if j.Raw.Status.Succeeded > 0 {
		console.Format("Succeeded: %d\n", j.Raw.Status.Succeeded)
	}

	console.Format("Job name: %s\n", j.Name)
	controller := string(j.Raw.UID)

	for _, p := range k.Pods(
		filter.New().AddNamespaces(namespace).AddNames(j.Name),
	) {
		if p.Label("batch.kubernetes.io/controller-uid") != controller {
			continue
		}

		console.Format("Pod: %s\n", p.Name)
		log := k.Log(namespace, p.Name, "")
		console.Line("Log:")
		console.Line(log)
		console.Line("-----")
	}
}
