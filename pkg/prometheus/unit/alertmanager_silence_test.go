package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestSilence(t *testing.T) {
	actual := silence.New(
		&models.GettableSilence{
			ID:     new(stringsConstant.UpperAlfa),
			Status: &models.SilenceStatus{State: new(constant.ActiveState)},
			Silence: models.Silence{
				CreatedBy: new(stringsConstant.UpperBravo),
				Comment:   new(stringsConstant.UpperCharlie),
				StartsAt:  new(time.Scan(assert.NewMinute(0))),
				EndsAt:    new(time.Scan(assert.NewMinute(10))),
			},
		},
		locator.New(stringsConstant.UpperDelta),
	)
	actual.Start = nil
	actual.End = nil
	actual.Raw = nil
	assert.Any(
		t,
		&silence.Silence{
			MonitorIdentifier: "silence-Alfa",
			Identifier:        "Alfa",
			State:             "active",
			Author:            "Bravo",
			Comment:           "Charlie",
			Rule:              "unknown rule",
			Link:              "https://Delta/#/silences/Alfa",
		},
		actual,
	)
}
