package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (s *Service) ListResources(
	x context.Context,
	clusterName string,
	q ListQuery,
) ([]response.ResourceSummary, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

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
		return nil, g
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
