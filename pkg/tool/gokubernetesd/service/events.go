package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
)

func (s *Service) Events(
	x context.Context,
	clusterName string,
	q EventsQuery,
) ([]response.EventEntry, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = v1.NamespaceAll
	}

	var selectors []string

	if q.Kind != "" {
		selectors = append(
			selectors,
			fmt.Sprintf("involvedObject.kind=%s", q.Kind),
		)
	}

	if q.Name != "" {
		selectors = append(
			selectors,
			fmt.Sprintf("involvedObject.name=%s", q.Name),
		)
	}

	if q.Type != "" {
		selectors = append(selectors, fmt.Sprintf("type=%s", q.Type))
	}

	events, f := c.Clientset().CoreV1().Events(namespace).List(
		x,
		v1.ListOptions{FieldSelector: join.Comma(selectors)},
	)

	if f != nil {
		return nil, f
	}

	sort.Slice(
		events.Items,
		func(i, j int) bool {
			return events.Items[i].LastTimestamp.After(events.Items[j].LastTimestamp.Time)
		},
	)
	limit := q.Limit

	if limit <= 0 {
		limit = 50
	}

	result := []response.EventEntry{}

	for _, ev := range events.Items {
		if len(result) >= limit {
			break
		}

		if !q.Unfiltered {
			muted, g := s.IsMuted(ev.Reason, ev.Message, clusterName)

			if g != nil {
				return nil, g
			}

			if muted {
				continue
			}
		}

		result = append(
			result,
			response.EventEntry{
				Namespace: ev.Namespace,
				Type:      ev.Type,
				Reason:    ev.Reason,
				Object: fmt.Sprintf(
					"%s/%s",
					ev.InvolvedObject.Kind,
					ev.InvolvedObject.Name,
				),
				Message: ev.Message,
				Count:   ev.Count,
				Age:     format.Age(ev.LastTimestamp.Time),
			},
		)
	}

	return result, nil
}
