package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestAlert(t *testing.T) {
	actual := alert.New(
		&models.GettableAlert{
			Fingerprint: new(strings.UpperAlfa),
			Status:      &models.AlertStatus{State: new(prometheus.ActiveState)},
			StartsAt:    new(strfmt.NewDateTime()),
		},
		strings.UpperBravo,
	)
	actual.Raw = nil
	assert.Exported(
		t,
		&alert.Alert{
			MonitorIdentifier: "alert-Alfa",
			Fingerprint:       "Alfa",
			Name:              "none",
			State:             "active",
			Severity:          "none",
			Summary:           "none",
			Message:           "none",
			Prometheus:        "none",
			Start:             new(libraryConstant.StartOfTime),
			Link:              "https://Bravo/#/alerts?filter=%7Balertname%3D%22none%22%7D",
		},
		actual,
	)
}
