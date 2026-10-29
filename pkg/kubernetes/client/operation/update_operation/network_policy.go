package update_operation

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
	np *networking.NetworkPolicy,
) error {
	_, e := operation.NetworkPolicy(c, np.Namespace).Update(
		x,
		np,
		meta.UpdateOptions{},
	)

	return refusal.Classify(e)
}
