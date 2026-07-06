package list

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	batch "k8s.io/api/batch/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Job(
	c *kubernetes.Clientset,
	x context.Context,
	namespace string,
	node string,
) []batch.Job {
	var o meta.ListOptions

	if node != constant.NodeAll {
		o.FieldSelector = fmt.Sprintf("spec.nodeName=%s", node)
	}

	r, e := operation.Job(c, namespace).List(x, o)
	errors.PanicOnError(e)

	return r.Items
}
