package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/nut"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"github.com/funtimecoding/soil/pkg/nut/detached"
	"github.com/funtimecoding/soil/pkg/nut/fixture"
	"testing"
)

func TestBatteriesListed(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{constant.ListBatteries: constant.FixtureBatteries},
	)
	defer s.Close()
	result, e := nut.New(s.Host(), s.Port()).Batteries(t.Context())
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"alfa", "bravo"}, result)
}

func TestVariablesUnquoted(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{"LIST VAR alfa": constant.FixtureAlfaVariables},
	)
	defer s.Close()
	result, e := nut.New(s.Host(), s.Port()).Variables(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "100", result["battery.charge"])
	assert.String(t, "OL CHRG", result["ups.status"])
	assert.String(t, `Charlie "Delta" 850 `, result["device.model"])
	assert.Count(t, 10, result)
}

func TestUnknownBatteryNotFound(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{"LIST VAR alfa": constant.FixtureUnknownBattery},
	)
	defer s.Close()
	_, e := nut.New(s.Host(), s.Port()).Variables(t.Context(), "alfa")
	assert.True(t, not_found.Is(e))
}

func TestStaleBatteryDetached(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{"LIST VAR alfa": constant.FixtureStaleReadings},
	)
	defer s.Close()
	_, e := nut.New(s.Host(), s.Port()).Variables(t.Context(), "alfa")
	assert.True(t, detached.Is(e))
}

func TestDriverlessBatteryDetached(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{"LIST VAR alfa": constant.FixtureNoDriver},
	)
	defer s.Close()
	_, e := nut.New(s.Host(), s.Port()).Variables(t.Context(), "alfa")
	assert.True(t, detached.Is(e))
}

func TestClosedServerUnreachable(t *testing.T) {
	s := fixture.NewServer(map[string]string{})
	host, port := s.Host(), s.Port()
	s.Close()
	assert.Error(t, nut.New(host, port).Reachable(t.Context()))
}
