package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestChanger(t *testing.T) {
	a1 := field_changer.New()
	a1.AddName("BigApple", "Apple")
	a1.AddName("ComplicatedOrange", "Orange")
	assert.String(t, "Apple", a1.Name("BigApple"))
	assert.String(t, "Orange", a1.Name("ComplicatedOrange"))
	assert.String(t, "Banana", a1.Name("Banana"))
	assert.Any(
		t,
		[]*alert.Alert{
			{Name: "Apple"},
			{Name: "Orange"},
		},
		a1.Run(
			[]*alert.Alert{
				{Name: "BigApple"},
				{Name: "ComplicatedOrange"},
			},
		),
	)
	a2 := field_changer.New()
	a2.AddName("BigApple", "Apple")
	a2.AddSeverity(
		"Apple",
		constant.WarningSeverity,
		constant.CriticalSeverity,
	)
	assert.Any(
		t,
		[]*alert.Alert{
			{Name: "Apple", Severity: "critical"},
		},
		a2.Run(
			[]*alert.Alert{
				{Name: "BigApple", Severity: constant.WarningSeverity},
			},
		),
	)
}
