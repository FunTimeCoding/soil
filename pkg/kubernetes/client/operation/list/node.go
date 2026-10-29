package list

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Node(
	c kubernetes.Interface,
	x context.Context,
) ([]core.Node, error) {
	r, e := operation.Node(c).List(x, meta.ListOptions{})

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return r.Items, nil
}
