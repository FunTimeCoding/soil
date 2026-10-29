package delete_operation

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Event(
	c kubernetes.Interface,
	x context.Context,
	namespace string,
	name string,
) error {
	return refusal.Classify(
		operation.Event(c, namespace).Delete(x, name, v1.DeleteOptions{}),
	)
}
