package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	nutConstant "github.com/funtimecoding/soil/pkg/nut/constant"
	"github.com/funtimecoding/soil/pkg/nut/fixture"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/unit/battery_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestBatteryMeasuresExported(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{
			nutConstant.ListBatteries: nutConstant.FixtureAlfaBattery,
			"LIST VAR alfa":           nutConstant.FixtureAlfaVariables,
		},
	)
	defer s.Close()
	b := battery_tester.New(t, battery_tester.Instance("golf", s))
	b.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			b.Metric.Registry(),
			strings.NewReader(
				`
# HELP battery_charge_percent Battery charge
# TYPE battery_charge_percent gauge
battery_charge_percent{battery="alfa",server="golf"} 100
# HELP battery_runtime_seconds Runtime left on battery at the current load
# TYPE battery_runtime_seconds gauge
battery_runtime_seconds{battery="alfa",server="golf"} 3061
# HELP battery_input_volts Input voltage
# TYPE battery_input_volts gauge
battery_input_volts{battery="alfa",server="golf"} 239
# HELP battery_status A battery's status flags, 1 for each flag it reports
# TYPE battery_status gauge
battery_status{battery="alfa",flag="charging",server="golf"} 1
battery_status{battery="alfa",flag="online",server="golf"} 1
# HELP battery_information A battery's model, manufacturer and serial, always 1
# TYPE battery_information gauge
battery_information{battery="alfa",manufacturer="Echo",model="Charlie \"Delta\" 850",serial="9B00",server="golf"} 1
# HELP battery_attached Whether the battery's driver serves its data, 1 or 0
# TYPE battery_attached gauge
battery_attached{battery="alfa",server="golf"} 1
# HELP battery_server_reachable Whether an inventory server answered the last poll, 1 or 0
# TYPE battery_server_reachable gauge
battery_server_reachable{server="golf"} 1
`,
			),
			"battery_charge_percent",
			"battery_runtime_seconds",
			"battery_input_volts",
			constant.StatusMetric,
			constant.InformationMetric,
			constant.AttachedMetric,
			constant.ReachableMetric,
		),
	)
}

func TestStaleBatteryDetachedQuietly(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{
			nutConstant.ListBatteries: nutConstant.FixtureAlfaBattery,
			"LIST VAR alfa":           nutConstant.FixtureStaleReadings,
		},
	)
	defer s.Close()
	b := battery_tester.New(t, battery_tester.Instance("golf", s))

	for range 2 {
		b.Poll()
	}

	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			b.Metric.Registry(),
			strings.NewReader(
				`
# HELP battery_attached Whether the battery's driver serves its data, 1 or 0
# TYPE battery_attached gauge
battery_attached{battery="alfa",server="golf"} 0
# HELP battery_server_reachable Whether an inventory server answered the last poll, 1 or 0
# TYPE battery_server_reachable gauge
battery_server_reachable{server="golf"} 1
`,
			),
			constant.AttachedMetric,
			constant.ReachableMetric,
		),
	)
	assert.Count(t, 0, b.Reporter.Events())
}

func TestUnreportedVariableNotExported(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{
			nutConstant.ListBatteries: nutConstant.FixtureAlfaBattery,
			"LIST VAR alfa":           nutConstant.FixtureAlfaVariables,
		},
	)
	defer s.Close()
	b := battery_tester.New(t, battery_tester.Instance("golf", s))
	b.Poll()
	count, e := testutil.GatherAndCount(
		b.Metric.Registry(),
		"battery_output_volts",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, count)
}

func TestUnknownStatusFlagPassesThrough(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{
			nutConstant.ListBatteries: nutConstant.FixtureBravoBattery,
			"LIST VAR bravo":          nutConstant.FixtureBravoVariables,
		},
	)
	defer s.Close()
	b := battery_tester.New(t, battery_tester.Instance("golf", s))
	b.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			b.Metric.Registry(),
			strings.NewReader(
				`
# HELP battery_status A battery's status flags, 1 for each flag it reports
# TYPE battery_status gauge
battery_status{battery="bravo",flag="FOXTROT",server="golf"} 1
battery_status{battery="bravo",flag="low_battery",server="golf"} 1
battery_status{battery="bravo",flag="on_battery",server="golf"} 1
`,
			),
			constant.StatusMetric,
		),
	)
}

func TestUnreachableServerQuiet(t *testing.T) {
	s := fixture.NewServer(map[string]string{})
	gone := battery_tester.Instance("hotel", s)
	s.Close()
	b := battery_tester.New(t, gone)

	for range 2 {
		b.Poll()
	}

	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			b.Metric.Registry(),
			strings.NewReader(
				`
# HELP battery_server_reachable Whether an inventory server answered the last poll, 1 or 0
# TYPE battery_server_reachable gauge
battery_server_reachable{server="hotel"} 0
`,
			),
			constant.ReachableMetric,
		),
	)
	assert.Count(t, 0, b.Reporter.Events())
}

func TestRefusedListCapturedOnce(t *testing.T) {
	s := fixture.NewServer(
		map[string]string{
			nutConstant.ListBatteries: nutConstant.FixtureAlfaBattery,
			"LIST VAR alfa":           nutConstant.FixtureAccessDenied,
		},
	)
	defer s.Close()
	b := battery_tester.New(t, battery_tester.Instance("golf", s))
	b.Poll()
	assert.Count(t, 1, b.Reporter.Events())
}
