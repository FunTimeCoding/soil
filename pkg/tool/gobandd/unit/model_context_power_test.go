package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestShutdownToolRefusesWithoutSendingAnything(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(band.PowerOn, hardStates()...)
	c := model_context_client.New(t, s.Port)
	defer c.Close()
	c.MustCallTool(
		constant.UseInstance,
		map[string]any{constant.InstanceParameter: "alfa"},
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
	c := model_context_client.New(t, s.Port)
	defer c.Close()
	c.MustCallTool(
		constant.UseInstance,
		map[string]any{constant.InstanceParameter: "alfa"},
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
