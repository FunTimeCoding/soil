package terraform_tester

import (
	"fmt"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	coordination "k8s.io/api/coordination/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func HeldLease(
	who string,
	created time.Time,
) *lease.Lease {
	return lease.New(
		&coordination.Lease{
			ObjectMeta: meta.ObjectMeta{
				Name:      constant.StateLease,
				Namespace: constant.StateNamespace,
				Annotations: map[string]string{
					constant.LockAnnotation: fmt.Sprintf(
						constant.FixtureLockTemplate,
						constant.FixtureLock,
						who,
						created.Format(time.RFC3339Nano),
					),
				},
			},
			Spec: coordination.LeaseSpec{
				HolderIdentity: new(constant.FixtureLock),
			},
		},
		kubernetesConstant.InCluster,
	)
}
