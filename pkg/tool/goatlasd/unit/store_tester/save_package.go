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

func SavePackage(
	t *testing.T,
	s *store.Store,
	name string,
	packageName string,
	version string,
	seenAt time.Time,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.SavePlacement(
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
