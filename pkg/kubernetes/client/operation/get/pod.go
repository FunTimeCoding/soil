package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Pod(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	name string,
) (*v1.Pod, error) {
	var o meta.GetOptions
	result, e := operation.Pod(c, namespace).Get(x, name, o)

	if e != nil {
		if errors.IsNotFound(e) {
			return nil, nil
		}

		return nil, e
	}

	return result, nil
}
