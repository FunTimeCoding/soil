package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	batch "k8s.io/api/batch/v1"
	kubernetesError "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CronJob(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) *batch.CronJob {
	var o meta.GetOptions
	result, e := operation.CronJob(c, namespace).Get(x, name, o)

	if e != nil {
		if kubernetesError.IsNotFound(e) {
			return nil
		}
	}

	errors.PanicOnError(e)

	return result
}
