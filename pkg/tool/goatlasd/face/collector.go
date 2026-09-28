package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

type Collector interface {
	Source() string
	Collect(
		q context.Context,
		s *gazetteer.Gazetteer,
	) ([]*placement.Placement, error)
}
