package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"testing"
)

func TestAnyPlaceFindsTheClaimedHost(t *testing.T) {
	assert.True(
		t,
		web.AnyPlace(
			[]*sighting.Sighting{
				collector_tester.NewSighting(
					constant.FixtureAddress,
					constant.FixtureHostname,
					place.New("", 0, ""),
				),
				collector_tester.NewSighting(
					constant.FixtureAddress,
					constant.FixtureHostname,
					place.New(
						constant.FixtureDeviceNode,
						constant.FixturePlaceIdentifier,
						constant.FixtureDeviceNode,
					),
				),
			},
		),
	)
}

func TestAnyPlaceRefusesAnUnclaimedList(t *testing.T) {
	assert.False(
		t,
		web.AnyPlace(
			[]*sighting.Sighting{
				collector_tester.NewSighting(
					constant.FixtureAddress,
					constant.FixtureHostname,
					place.New("", 0, ""),
				),
			},
		),
	)
}
