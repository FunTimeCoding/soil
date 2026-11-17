package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"testing"
)

func TestUpdateDeviceTool(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	s.Source.SeedDeviceType("bravo")
	var d client.Device
	notation.MustDecode(
		s.Tools().MustCallTool(
			constant.UpdateDevice,
			map[string]any{
				constant.Device:       "alfa",
				constant.Type:         "bravo",
				constant.LocalContext: map[string]any{"charlie": "delta"},
			},
		),
		&d,
		false,
	)
	assert.Any(t, new("bravo"), d.Type)
	assert.Any(t, &map[string]any{"charlie": "delta"}, d.LocalContext)
}

func TestUpdateDeviceToolEmptyLocalContextClears(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	c := s.Tools()
	set := c.MustCallToolNotation(
		constant.UpdateDevice,
		map[string]any{
			constant.Device:       "alfa",
			constant.LocalContext: map[string]any{"charlie": "delta"},
		},
	)
	assert.Any(t, map[string]any{"charlie": "delta"}, set["localContext"])
	result := c.MustCallToolNotation(
		constant.UpdateDevice,
		map[string]any{
			constant.Device:       "alfa",
			constant.LocalContext: map[string]any{},
		},
	)
	assert.Nil(t, result["localContext"])
}

func TestUpdateVirtualMachineTool(t *testing.T) {
	s := base.New(t)
	s.Source.SeedVirtualMachine("alfa")
	result := s.Tools().MustCallToolNotation(
		constant.UpdateVirtualMachine,
		map[string]any{
			constant.VirtualMachine: "alfa",
			constant.LocalContext:   map[string]any{"bravo": "charlie"},
		},
	)
	assert.Any(t, map[string]any{"bravo": "charlie"}, result["localContext"])
}

func TestUpdateDeviceToolWithEmptyDeviceFails(t *testing.T) {
	s := base.New(t)
	message := s.Tools().MustCallToolError(
		constant.UpdateDevice,
		map[string]any{constant.Device: "", constant.Type: "bravo"},
	)
	assert.StringContains(t, "device is required", message)
}
