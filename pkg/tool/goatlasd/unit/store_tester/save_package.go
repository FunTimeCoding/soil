package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func (o *Tester) SavePackage(
	name string,
	packageName string,
	version string,
	seenAt time.Time,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.SavePlacement(
			placement.NewPackage(
				constant.SourceOutpost,
				constant.KindService,
				constant.FixtureScope,
				name,
				packageName,
				version,
				place.New(
					constant.FixtureDeviceNode,
					1,
					constant.FixtureDeviceNode,
				),
				seenAt,
			),
		),
	)
}
