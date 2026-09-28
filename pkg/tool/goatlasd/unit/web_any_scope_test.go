package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"testing"
)

func TestAnyScopeFindsTheScopedRow(t *testing.T) {
	assert.True(
		t,
		web.AnyScope(
			[]*placement.Placement{
				collector_tester.NewScopelessPlacement(
					constant.SourceProcess,
					constant.FixtureService,
				),
				collector_tester.NewPlacement(
					constant.SourceKubernetes,
					constant.FixturePeerService,
				),
			},
		),
	)
}

func TestAnyScopeRefusesAnUnscopedSource(t *testing.T) {
	assert.False(
		t,
		web.AnyScope(
			[]*placement.Placement{
				collector_tester.NewScopelessPlacement(
					constant.SourceProcess,
					constant.FixtureService,
				),
			},
		),
	)
}
