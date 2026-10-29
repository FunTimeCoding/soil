package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/exec_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
)

func (s *Service) ExecInPod(
	x context.Context,
	clusterName string,
	q request.Exec,
) (*exec_result.Result, error) {
	if q.Name == "" {
		return nil, validation.New("name is required")
	}

	if q.Namespace == "" {
		return nil, validation.New("namespace is required")
	}

	if len(q.Command) == 0 {
		return nil, validation.New("command is required")
	}

	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*exec_result.Result, error) {
			return execInPodOn(x, c, q)
		},
	)
}
