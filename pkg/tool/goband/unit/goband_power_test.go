package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goband/plan"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestDryRunSaysWhenShutdownWouldBeRefused(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerCycle,
		bandConstant.PowerOffSoft,
	)
	output, _, code := run(t, s, "shutdown", "--instance", "alfa")
	assert.Integer(t, 0, code)
	var p plan.Power
	notation.MustDecode(output, &p, false)
	assert.String(t, "alfa", p.Instance)
	assert.NotNil(t, p.Refusal)
	assert.String(
		t,
		"off soft graceful not offered: LMS on the host registers it; force cuts power instead",
		*p.Refusal,
	)
	assert.String(t, "", p.Description)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestDryRunDescribesAcceptedShutdown(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerOffSoft,
		bandConstant.PowerOffSoftGraceful,
	)
	output, _, code := run(t, s, "shutdown", "--instance", "alfa")
	assert.Integer(t, 0, code)
	var p plan.Power
	notation.MustDecode(output, &p, false)
	assert.String(t, "alfa", p.Instance)
	assert.Nil(t, p.Refusal)
	assert.String(
		t,
		"Would ask the operating system to shut down through LMS. Pass --yes to execute, then poll status - if the machine never goes down, the operating system is not answering and --force is the remaining option.",
		p.Description,
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
	var r client.PowerResult
	notation.MustDecode(output, &r, false)
	assert.String(t, "alfa", r.Instance)
	assert.String(t, "off soft graceful", r.Requested)
	assert.String(t, "on", r.Power)
	assert.Integers(t, []int{12}, s.Alfa.Requested())
}

func TestRefusedShutdownExitsWithDaemonsMessage(t *testing.T) {
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

func TestStartWithoutInstanceIsRefusedWhileTwoAreConfigured(t *testing.T) {
	s := base.New(t)
	_, failure, code := run(t, s, "start", "--yes")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "selection required", failure)
	assert.Count(t, 0, s.Alfa.Requested())
	assert.Count(t, 0, s.Bravo.Requested())
}
