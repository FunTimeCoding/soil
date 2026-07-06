package delete_operation

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Event(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) {
	errors.PanicOnError(
		operation.Event(c, namespace).Delete(
			x,
			name,
			v1.DeleteOptions{},
		),
	)
}
