package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"testing"
)

func TestAnyScopeFindsScopedRow(t *testing.T) {
	assert.True(
		t,
		component.AnyScope(
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

func TestAnyScopeRefusesUnscopedSource(t *testing.T) {
	assert.False(
		t,
		component.AnyScope(
			[]*placement.Placement{
				collector_tester.NewScopelessPlacement(
					constant.SourceProcess,
					constant.FixtureService,
				),
			},
		),
	)
}
