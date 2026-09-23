package unit

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"k8s.io/api/coordination/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func annotated(v map[string]string) *lease.Lease {
	return lease.New(
		&v1.Lease{ObjectMeta: v11.ObjectMeta{Annotations: v}},
		"in-cluster",
	)
}
