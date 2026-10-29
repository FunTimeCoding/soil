package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Pod(
	c kubernetes.Interface,
	x context.Context,
	namespace string,
	name string,
) (*v1.Pod, error) {
	result, e := operation.Pod(c, namespace).Get(x, name, meta.GetOptions{})

	if e != nil {
		if errors.IsNotFound(e) {
			return nil, nil
		}

		return nil, refusal.Classify(e)
	}

	return result, nil
}
