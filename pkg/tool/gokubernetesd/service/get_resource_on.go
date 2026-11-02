package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/filter_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func getResourceOn(
	x context.Context,
	c *cluster.Cluster,
	q request.Get,
) (*filter_result.FilterResult, error) {
	gvr, namespaced, f := c.ResolveResource(q.ResourceType)

	if f != nil {
		return nil, f
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	var result *unstructured.Unstructured
	var g error

	if namespaced {
		result, g = c.Dynamic().Resource(gvr).Namespace(namespace).Get(
			x,
			q.Name,
			v1.GetOptions{},
		)
	} else {
		result, g = c.Dynamic().Resource(gvr).Get(x, q.Name, v1.GetOptions{})
	}

	if g != nil {
		return nil, refusal.Classify(g)
	}

	return resource.FilterObject(result.Object, q.Unfiltered), nil
}
