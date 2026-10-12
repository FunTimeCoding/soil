package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestStatusReportsPowerAndTheOperationsOffered(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(
		constant.PowerOn,
		constant.PowerCycle,
		constant.PowerOffSoft,
		constant.PowerOffSoftGraceful,
	)
	r, e := s.Client.GetStatusWithResponse(
		t.Context(),
		&client.GetStatusParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.String(t, "on", r.JSON200.Power)
	assert.Strings(
		t,
		[]string{"power cycle", "off soft", "off soft graceful"},
		r.JSON200.Available,
	)
}
