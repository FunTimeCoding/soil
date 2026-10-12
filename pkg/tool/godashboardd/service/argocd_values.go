package service

import (
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"strconv"
)

func (s *Service) argocdValues() []string {
	pending := []string{
		constant.PendingValue,
		constant.PendingValue,
		constant.PendingValue,
		constant.PendingValue,
	}

	if s.argocd == nil {
		return pending
	}

	applications, e := s.argocd.Applications()

	if e != nil {
		s.logger.Plain("applications fetch failed: %v", e)

		return pending
	}

	var outOfSync, degraded, missing int

	for _, v := range applications {
		if v.Sync != kubernetesConstant.ArgocdSyncSynced {
			outOfSync++
		}

		if v.Health == kubernetesConstant.ArgocdHealthDegraded {
			degraded++
		}

		if v.Health == kubernetesConstant.ArgocdHealthMissing {
			missing++
		}
	}

	return []string{
		strconv.Itoa(len(applications)),
		strconv.Itoa(outOfSync),
		strconv.Itoa(degraded),
		strconv.Itoa(missing),
	}
}
