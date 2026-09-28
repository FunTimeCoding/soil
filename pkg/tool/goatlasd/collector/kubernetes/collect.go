package kubernetes

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (c *Collector) Collect(
	_ context.Context,
	s *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	var result []*placement.Placement
	seen := map[key]bool{}
	now := time.Now()

	for _, p := range c.kubernetes.Pods(nil) {
		node := p.Raw.Spec.NodeName
		k := key{scope: p.Raw.Namespace, name: ApplicationName(p), node: node}

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
				k.scope,
				k.name,
				where,
				now,
			),
		)
	}

	return result, nil
}
