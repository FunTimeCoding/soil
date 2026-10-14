package runner

import (
	"github.com/funtimecoding/soil/pkg/kubernetes"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/lock_detail"
)

func (r *Runner) orphaned(d *lock_detail.Detail) bool {
	host := d.Host()

	if host == "" || r.hostname == "" {
		return false
	}

	if host == r.hostname {
		return d.Created.Before(r.started)
	}

	deployment := kubernetes.PodDeployment(host)

	if deployment == "" || deployment != kubernetes.PodDeployment(r.hostname) {
		return false
	}

	if r.namespace == "" || len(host) >= kubernetesConstant.HostnameLimit {
		return false
	}

	p, e := r.kubernetes.Pod(r.namespace, host)

	if e != nil {
		r.reporter.CaptureWithContext(
			e,
			constant.LockContext,
			map[string]any{"lock": d.Identifier, "who": d.Who},
		)

		return false
	}

	return p == nil
}
