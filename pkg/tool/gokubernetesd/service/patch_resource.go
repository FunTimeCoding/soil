package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) PatchResource(
	x context.Context,
	clusterName string,
	q request.Patch,
) error {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return e
	}

	gvr, namespaced, f := c.ResolveResource(q.ResourceType)

	if f != nil {
		return f
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	options := v1.PatchOptions{}

	if q.DryRun {
		options.DryRun = []string{v1.DryRunAll}
	}

	if namespaced {
		_, g := c.Dynamic().Resource(gvr).Namespace(namespace).Patch(
			x,
			q.Name,
			q.Type,
			[]byte(q.Patch),
			options,
		)

		return g
	}

	_, g := c.Dynamic().Resource(gvr).Patch(
		x,
		q.Name,
		q.Type,
		[]byte(q.Patch),
		options,
	)

	return g
}
