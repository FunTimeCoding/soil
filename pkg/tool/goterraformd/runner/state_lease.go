package runner

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"

func (r *Runner) stateLease() *lease.Lease {
	result, e := r.kubernetes.Lease(r.stateNamespace, r.stateLeaseName)

	if e != nil {
		return nil
	}

	return result
}
