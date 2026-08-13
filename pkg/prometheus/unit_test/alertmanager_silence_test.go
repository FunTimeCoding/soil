package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestSilence(t *testing.T) {
	actual := silence.New(
		&models.GettableSilence{
			ID:     new(strings.UpperAlfa),
			Status: &models.SilenceStatus{State: new(prometheus.ActiveState)},
			Silence: models.Silence{
				CreatedBy: new(strings.UpperBravo),
				Comment:   new(strings.UpperCharlie),
				StartsAt:  new(time.Scan(assert.NewMinute(0))),
				EndsAt:    new(time.Scan(assert.NewMinute(10))),
			},
		},
		strings.UpperDelta,
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
