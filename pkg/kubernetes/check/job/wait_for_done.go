package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"time"
)

func waitForDone(
	k *client.Client,
	namespace string,
	job string,
) {
	console.Line("Sleep before wait")
	time.Sleep(10 * time.Second)
	printJobs(k, namespace)
	console.Line("Wait for job")
	errors.PanicOnError(k.WaitForJob(namespace, job, 0))
}
