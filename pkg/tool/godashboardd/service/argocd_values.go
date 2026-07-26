package service

import (
	argocd "github.com/funtimecoding/soil/pkg/argocd/constant"
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
		if v.Sync != argocd.Synced {
			outOfSync++
		}

		if v.Health == argocd.Degraded {
			degraded++
		}

		if v.Health == argocd.Missing {
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
