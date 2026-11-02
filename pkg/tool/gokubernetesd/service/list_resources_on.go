package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func listResourcesOn(
	x context.Context,
	c *cluster.Cluster,
	q request.List,
) ([]response.ResourceSummary, error) {
	gvr, namespaced, f := c.ResolveResource(q.ResourceType)

	if f != nil {
		return nil, f
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	options := v1.ListOptions{}

	if q.LabelSelector != "" {
		options.LabelSelector = q.LabelSelector
	}

	if q.FieldSelector != "" {
		options.FieldSelector = q.FieldSelector
	}

	var list *unstructured.UnstructuredList
	var g error

	if !namespaced || q.AllNamespaces {
		list, g = c.Dynamic().Resource(gvr).List(x, options)
	} else {
		list, g = c.Dynamic().Resource(gvr).Namespace(namespace).List(
			x,
			options,
		)
	}

	if g != nil {
		return nil, refusal.Classify(g)
	}

	result := []response.ResourceSummary{}

	for _, u := range list.Items {
		i := response.ResourceSummary{
			Name:      u.GetName(),
			Namespace: u.GetNamespace(),
			Kind:      u.GetKind(),
			Age:       format.Age(u.GetCreationTimestamp().Time),
		}
		i.Status = resource.ExtractStatus(&u)
		i.Restarts = resource.ExtractRestarts(&u)
		result = append(result, i)
	}

	return result, nil
}
