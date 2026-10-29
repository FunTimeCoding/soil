package get

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	networking "k8s.io/api/networking/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NetworkPolicy(
	c kubernetes.Interface,
	x context.Context,
	namespace string,
	name string,
) (*networking.NetworkPolicy, error) {
	result, e := operation.NetworkPolicy(c, namespace).Get(
		x,
		name,
		meta.GetOptions{},
	)

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return result, nil
}
