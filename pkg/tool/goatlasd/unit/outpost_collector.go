package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
)

func outpostCollector(sources ...face.OutpostSource) *outpost.Collector {
	var targets []*target.Target

	for i, s := range sources {
		targets = append(
			targets,
			target.New(constant.FixtureOutpostNames[i], s),
		)
	}

	return outpost.New(targets, logger.New(context.Background()))
}
