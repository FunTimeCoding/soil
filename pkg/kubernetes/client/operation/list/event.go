package list

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	event "k8s.io/api/events/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Event(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	limit int,
	selector string,
) []event.Event {
	r, e := operation.Event(c, namespace).List(
		x,
		meta.ListOptions{Limit: int64(limit), FieldSelector: selector},
	)
	errors.PanicOnError(e)

	return r.Items
}
