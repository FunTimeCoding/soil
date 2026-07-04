package service

import (
	"github.com/funtimecoding/go-library/pkg/argocd/application"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
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
		if v.Sync != application.Synced {
			outOfSync++
		}

		if v.Health == application.Degraded {
			degraded++
		}

		if v.Health == application.Missing {
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
