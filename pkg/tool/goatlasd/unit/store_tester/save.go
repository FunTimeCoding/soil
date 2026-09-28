package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"testing"
	"time"
)

func Save(
	t *testing.T,
	s *store.Store,
	source string,
	name string,
	seenAt time.Time,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.SavePlacement(
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
