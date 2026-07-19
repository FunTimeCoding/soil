package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestSilence(t *testing.T) {
	actual := silence.New(
		&models.GettableSilence{
			ID: new(upper.Alfa),
			Status: &models.SilenceStatus{
				State: new(constant.ActiveState),
			},
			Silence: models.Silence{
				CreatedBy: new(upper.Bravo),
				Comment:   new(upper.Charlie),
				StartsAt:  new(time.Scan(assert.NewMinute(0))),
				EndsAt:    new(time.Scan(assert.NewMinute(10))),
			},
		},
		upper.Delta,
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
