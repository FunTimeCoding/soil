package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"strconv"
)

func (s *Service) argocdValues() ([]string, error) {
	pending := []string{
		constant.PendingValue,
		constant.PendingValue,
		constant.PendingValue,
		constant.PendingValue,
	}

	if s.argocd == nil {
		return pending, nil
	}

	applications, e := peer.Guarded(
		context.Background(),
		s.argocdGuard,
		func(context.Context) ([]*application.Application, error) {
			return s.argocd.Applications()
		},
	)

	if e != nil {
		return pending, e
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
	}, nil
}
