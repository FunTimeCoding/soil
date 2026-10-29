package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	coordination "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Lease(
	c kubernetes.Interface,
	x context.Context,
	namespace string,
	name string,
) (*coordination.Lease, error) {
	result, e := operation.Lease(c, namespace).Get(x, name, meta.GetOptions{})

	if e != nil {
		if errors.IsNotFound(e) {
			return nil, nil
		}

		return nil, refusal.Classify(e)
	}

	return result, nil
}
