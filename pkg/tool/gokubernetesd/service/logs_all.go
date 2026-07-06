package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
)

func (s *Service) LogsAll(
	x context.Context,
	clusterName string,
	q LogsQuery,
) ([]response.PodLog, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	pods, f := ResolvePods(x, c, q.Name, q.Namespace)

	if f != nil {
		return nil, f
	}

	if len(pods) == 0 {
		return nil, fmt.Errorf(
			"no pods found for %s in %s",
			q.Name,
			q.Namespace,
		)
	}

	result := []response.PodLog{}

	for _, pod := range pods {
		content, g := PodLogsString(
			x,
			c,
			pod,
			q.Namespace,
			q.Container,
			q.Tail,
			q.Since,
			q.Previous,
			q.Timestamps,
		)

		if g != nil {
			result = append(
				result,
				response.PodLog{
					Pod:     pod,
					Content: fmt.Sprintf("error: %s", g),
				},
			)

			continue
		}

		result = append(
			result,
			response.PodLog{
				Pod:     pod,
				Content: content,
			},
		)
	}

	return result, nil
}
