package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/describe_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (s *Service) DescribeResource(
	x context.Context,
	clusterName string,
	q DescribeQuery,
) (*describe_result.Result, error) {
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

	var object *unstructured.Unstructured
	var g error

	if namespaced {
		object, g = c.Dynamic().Resource(gvr).Namespace(namespace).Get(
			x,
			q.Name,
			v1.GetOptions{},
		)
	} else {
		object, g = c.Dynamic().Resource(gvr).Get(
			x,
			q.Name,
			v1.GetOptions{},
		)
	}

	if g != nil {
		return nil, g
	}

	selectors := []string{
		fmt.Sprintf("involvedObject.name=%s", q.Name),
	}

	if namespaced {
		selectors = append(
			selectors,
			fmt.Sprintf("involvedObject.namespace=%s", namespace),
		)
	}

	events, h := c.Clientset().CoreV1().Events(namespace).List(
		x,
		v1.ListOptions{FieldSelector: join.Comma(selectors)},
	)
	eventList := []response.DescribeEvent{}

	if h == nil {
		for _, ev := range events.Items {
			eventList = append(
				eventList,
				response.DescribeEvent{
					Type:    ev.Type,
					Reason:  ev.Reason,
					Message: ev.Message,
					Age:     format.Age(ev.LastTimestamp.Time),
					Count:   ev.Count,
				},
			)
		}
	}

	filtered := resource.FilterObject(object.Object, q.Unfiltered)

	return describe_result.New(
		filtered.Object,
		eventList,
		filtered.Filtered,
	), nil
}
