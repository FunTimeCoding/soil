package outpost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (c *Collector) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	var result []*placement.Placement
	var failed []string

	for _, t := range c.targets {
		v, e := collectOne(q, s, t.Source)

		if e != nil {
			failed = append(failed, t.Name)
			c.logger.Plain(constant.OutpostTargetFailedFormat, t.Name, e)

			continue
		}

		result = append(result, v...)
	}

	if len(c.targets) > 0 && len(failed) == len(c.targets) {
		return nil, unexpected.Format(
			constant.OutpostEveryTargetFailed,
			join.Comma(failed),
		)
	}

	return result, nil
}
