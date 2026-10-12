package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestStatusSaysWhyAGracefulShutdownWouldBeRefused(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(constant.PowerOn, hardStates()...)
	r, e := s.Client.GetStatusWithResponse(
		t.Context(),
		&client.GetStatusParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 5, r.JSON200.Operations)

	for _, o := range r.JSON200.Operations {
		if string(o.Name) != "shutdown" || o.Force {
			continue
		}

		assert.String(t, "off soft graceful", o.State)
		assert.NotNil(t, o.Refusal)
		assert.StringContains(t, "not offered", *o.Refusal)
	}
}

func TestStatusLeavesAnAcceptedForcedShutdownWithoutRefusal(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(constant.PowerOn, hardStates()...)
	r, e := s.Client.GetStatusWithResponse(
		t.Context(),
		&client.GetStatusParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)

	for _, o := range r.JSON200.Operations {
		if string(o.Name) == "shutdown" && o.Force {
			assert.String(t, "off soft", o.State)
			assert.True(t, o.Refusal == nil)
		}
	}
}
