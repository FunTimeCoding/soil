package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func NewPlacement(
	source string,
	name string,
) *placement.Placement {
	return placement.New(
		source,
		constant.KindService,
		constant.FixtureScope,
		name,
		place.New(
			constant.FixtureDeviceNode,
			constant.FixturePlaceIdentifier,
			constant.FixtureDeviceNode,
		),
		time.Now(),
	)
}
