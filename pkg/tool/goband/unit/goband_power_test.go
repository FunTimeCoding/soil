package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestDryRunSaysWhenAShutdownWouldBeRefused(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerCycle,
		bandConstant.PowerOffSoft,
	)
	output, _, code := run(t, s, "shutdown", "--instance", "alfa")
	assert.Integer(t, 0, code)
	assert.StringContains(
		t,
		"Would be refused: off soft graceful not offered",
		output,
	)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestDryRunDescribesAnAcceptedShutdown(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerOffSoft,
		bandConstant.PowerOffSoftGraceful,
	)
	output, _, code := run(t, s, "shutdown", "--instance", "alfa")
	assert.Integer(t, 0, code)
	assert.StringContains(
		t,
		"Would ask the operating system to shut down through LMS. Pass --yes to execute.",
		output,
	)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestShutdownWithYesRequestsOffSoftGraceful(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerOffSoft,
		bandConstant.PowerOffSoftGraceful,
	)
	output, _, code := run(t, s, "shutdown", "--instance", "alfa", "--yes")
	assert.Integer(t, 0, code)
	assert.StringContains(t, "off soft graceful requested on alfa", output)
	assert.Integers(t, []int{12}, s.Alfa.Requested())
}

func TestARefusedShutdownExitsWithTheDaemonsMessage(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerCycle,
		bandConstant.PowerOffSoft,
	)
	_, failure, code := run(t, s, "shutdown", "--instance", "alfa", "--yes")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "power: off soft graceful not offered", failure)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestStartWithoutAnInstanceIsRefusedWhileTwoAreConfigured(t *testing.T) {
	s := base.New(t)
	_, failure, code := run(t, s, "start", "--yes")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "selection required", failure)
	assert.Count(t, 0, s.Alfa.Requested())
	assert.Count(t, 0, s.Bravo.Requested())
}
