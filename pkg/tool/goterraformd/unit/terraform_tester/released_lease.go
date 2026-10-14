package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"time"
)

func ReleasedLease(
	who string,
	created time.Time,
) *lease.Lease {
	result := HeldLease(who, created)
	result.Raw.Spec.HolderIdentity = nil

	return result
}
