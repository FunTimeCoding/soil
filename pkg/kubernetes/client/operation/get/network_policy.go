package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	networking "k8s.io/api/networking/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NetworkPolicy(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) *networking.NetworkPolicy {
	result, e := operation.NetworkPolicy(c, namespace).Get(
		x,
		name,
		meta.GetOptions{},
	)
	errors.PanicOnError(e)

	return result
}
