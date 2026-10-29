package kubernetes

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/kubernetes_key"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	var pods []*pod.Pod
	e := c.guard.Attempt(
		q,
		func(context.Context) error {
			var f error
			pods, f = c.kubernetes.Pods(nil)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	var result []*placement.Placement
	seen := map[kubernetes_key.Key]bool{}
	now := time.Now()

	for _, p := range pods {
		node := p.Raw.Spec.NodeName
		k := kubernetes_key.Key{
			Scope: p.Raw.Namespace,
			Name:  ApplicationName(p),
			Node:  node,
		}

		if seen[k] {
			continue
		}

		seen[k] = true
		where, _ := s.Resolve(node)
		result = append(
			result,
			placement.New(
				constant.SourceKubernetes,
				constant.KindService,
				k.Scope,
				k.Name,
				where,
				now,
			),
		)
	}

	return result, nil
}
