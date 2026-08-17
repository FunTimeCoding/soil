package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	coordination "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Lease(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) (*coordination.Lease, error) {
	var o meta.GetOptions
	result, e := operation.Lease(c, namespace).Get(x, name, o)

	if e != nil {
		if errors.IsNotFound(e) {
			return nil, nil
		}

		return nil, e
	}

	return result, nil
}
