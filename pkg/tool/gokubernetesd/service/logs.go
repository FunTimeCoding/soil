package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/ambiguous_pods"
)

func (s *Service) Logs(
	x context.Context,
	clusterName string,
	q LogsQuery,
) (string, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return "", e
	}

	pods, f := ResolvePods(x, c, q.Name, q.Namespace)

	if f != nil {
		return "", f
	}

	if len(pods) == 0 {
		return "", not_found.Format(
			"no pods found for %s in %s",
			q.Name,
			q.Namespace,
		)
	}

	if len(pods) > 1 {
		return "", ambiguous_pods.New(q.Name, pods)
	}

	return PodLogsString(
		x,
		c,
		pods[0],
		q.Namespace,
		q.Container,
		q.Tail,
		q.Since,
		q.Previous,
		q.Timestamps,
	)
}
