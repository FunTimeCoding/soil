package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestAlert(t *testing.T) {
	actual := alert.New(
		&models.GettableAlert{
			Fingerprint: new(stringsConstant.UpperAlfa),
			Status:      &models.AlertStatus{State: new(constant.ActiveState)},
			StartsAt:    new(strfmt.NewDateTime()),
		},
		locator.New(stringsConstant.UpperBravo),
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
			Start:             new(soil.StartOfTime),
			Link:              "https://Bravo/#/alerts?filter=%7Balertname%3D%22none%22%7D",
		},
		actual,
	)
}
