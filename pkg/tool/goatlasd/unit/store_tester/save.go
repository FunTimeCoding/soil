package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func (o *Tester) Save(
	source string,
	name string,
	seenAt time.Time,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.SavePlacement(
			placement.New(
				source,
				constant.KindService,
				constant.FixtureScope,
				name,
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
