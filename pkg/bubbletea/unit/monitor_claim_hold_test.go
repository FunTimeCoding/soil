package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"testing"
)

func TestAWindowReceivesClaimsThroughTheHold(t *testing.T) {
	s := base.New(t)
	g := nextClaimMessage(
		t,
		s.Port,
		generativeConstant.ModelContextTestToken,
		memory.New(),
	)
	assert.Nil(t, g.Error)
}

func TestAWindowShowsARefusedClaimStream(t *testing.T) {
	s := base.New(t)
	r := memory.New()
	g := nextClaimMessage(t, s.Port, "bogus", r)
	assert.StringContains(t, "stream status: 401", g.Error.Error())
	assert.Count(t, 0, r.Events())
}
