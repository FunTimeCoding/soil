package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
)

func deleteJob(
	k *client.Client,
	namespace string,
	name string,
) {
	console.Format("Delete job %s in %s\n", name, namespace)
	j, e := k.Job(namespace, name)

	if not_found.Is(e) {
		return
	}

	errors.PanicOnError(e)
	console.Format("  %s\n", j.Name)
	k.MustDeleteJobWatch(namespace, j.Name)
}
