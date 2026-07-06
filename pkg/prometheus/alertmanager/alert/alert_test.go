package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/linux/systemd/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestAlert(t *testing.T) {
	actual := New(
		&models.GettableAlert{
			Fingerprint: new(upper.Alfa),
			Status: &models.AlertStatus{
				State: new(constant.ActiveState),
			},
			StartsAt: new(strfmt.NewDateTime()),
		},
		upper.Bravo,
	)
	actual.instance = nil
	actual.Raw = nil
	assert.Any(
		t,
		&Alert{
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
