package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	band "github.com/funtimecoding/soil/pkg/band/constant"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestShutdownToolRefusesWithoutSendingAnything(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(band.PowerOn, hardStates()...)
	c := s.Connect(t)
	defer c.Close()
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)
	assert.StringContains(
		t,
		"off soft graceful not offered",
		c.MustCallToolError(constant.Shutdown, nil),
	)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestShutdownToolWithForceCutsPower(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(band.PowerOn, runningStates()...)
	c := s.Connect(t)
	defer c.Close()
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)
	assert.String(
		t,
		"off soft requested on alfa",
		c.MustCallTool(
			constant.Shutdown,
			map[string]any{constant.ForceParameter: true},
		),
	)
	assert.Integers(t, []int{8}, s.Alfa.Requested())
}
