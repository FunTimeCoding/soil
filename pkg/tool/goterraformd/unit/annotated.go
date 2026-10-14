package unit

import (
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	coordination "k8s.io/api/coordination/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func annotated(v map[string]string) *lease.Lease {
	return lease.New(
		&coordination.Lease{ObjectMeta: meta.ObjectMeta{Annotations: v}},
		kubernetesConstant.InCluster,
	)
}
