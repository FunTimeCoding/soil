package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"testing"
)

func TestPowerToolsShowConnection(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	s.Source.SeedDevice("bravo")
	c := s.Tools()
	c.MustCallToolNotation(
		constant.CreatePowerOutlet,
		map[string]any{constant.Device: "alfa", "name": "charlie"},
	)
	c.MustCallToolNotation(
		constant.CreatePowerPort,
		map[string]any{constant.Device: "bravo", "name": "delta"},
	)
	c.MustCallToolNotation(
		constant.CreatePowerCable,
		map[string]any{
			constant.OutletDevice: "alfa",
			constant.Outlet:       "charlie",
			constant.PortDevice:   "bravo",
			constant.Port:         "delta",
		},
	)
	var result []map[string]any
	notation.MustDecode(
		c.MustCallTool(
			constant.ListPowerOutlets,
			map[string]any{constant.Device: "alfa"},
		),
		&result,
		false,
	)
	assert.Count(t, 1, result)
	assert.Any(t, "bravo", result[0]["connectedDevice"])
	assert.Any(t, "delta", result[0]["connectedPort"])
}

func TestCreateCableToolSidesInOrder(t *testing.T) {
	s := base.New(t)
	s.Source.SeedInterface(s.Source.SeedDevice("alfa"), "bravo")
	s.Source.SeedInterface(s.Source.SeedDevice("charlie"), "delta")
	s.Tools().MustCallToolNotation(
		constant.CreateCable,
		map[string]any{
			constant.DeviceA:    "alfa",
			constant.InterfaceA: "bravo",
			constant.DeviceB:    "charlie",
			constant.InterfaceB: "delta",
		},
	)
	assert.Contains(t, []string{"CreateCable bravo delta"}, s.Source.Calls())
}

func TestEthernetPowerToolModeAndType(t *testing.T) {
	s := base.New(t)
	s.Source.SeedInterface(s.Source.SeedDevice("alfa"), "bravo")
	result := s.Tools().MustCallToolNotation(
		constant.SetInterfaceEthernetPower,
		map[string]any{
			constant.Device:    "alfa",
			constant.Interface: "bravo",
			constant.Mode:      "pd",
			constant.Type:      "type1-ieee802.3af",
		},
	)
	assert.Any(t, "pd", result["ethernetPowerMode"])
	assert.Any(t, "type1-ieee802.3af", result["ethernetPowerType"])
}

func TestCreateAddressRangeToolArgumentsInOrder(t *testing.T) {
	s := base.New(t)
	s.Tools().MustCallToolNotation(
		constant.CreateAddressRange,
		map[string]any{
			constant.Start:       "192.0.2.10/24",
			constant.End:         "192.0.2.20/24",
			constant.Status:      "reserved",
			constant.Description: "echo",
		},
	)
	assert.Contains(
		t,
		[]string{
			"CreateInternetAddressRange 192.0.2.10/24 192.0.2.20/24 reserved echo",
		},
		s.Source.Calls(),
	)
}
