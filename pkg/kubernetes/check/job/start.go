package job

import "github.com/funtimecoding/soil/pkg/kubernetes/client"

func start(
	k *client.Client,
	wait bool,
	namespace string,
	cron string,
	name string,
) {
	deleteJob(k, namespace, name)
	j := createJob(k, namespace, cron, name)

	if wait {
		waitForDone(k, namespace, j.Name)
		printLog(k, namespace, j.Name)
	}
}
