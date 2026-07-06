package list

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Namespace(
	c *kubernetes.Clientset,
	x context.Context,
) []core.Namespace {
	r, e := operation.Namespace(c).List(x, meta.ListOptions{})
	errors.PanicOnError(e)

	return r.Items
}
