package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func patchResourceOn(
	x context.Context,
	c *cluster.Cluster,
	q request.Patch,
) error {
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

		return refusal.Classify(g)
	}

	_, g := c.Dynamic().Resource(gvr).Patch(
		x,
		q.Name,
		q.Type,
		[]byte(q.Patch),
		options,
	)

	return refusal.Classify(g)
}
