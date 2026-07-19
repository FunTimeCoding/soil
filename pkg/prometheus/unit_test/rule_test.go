package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
)

func TestRule(t *testing.T) {
	actualAlert := rule.NewAlert(
		&v1.AlertingRule{Name: upper.Alfa},
		&v1.RuleGroup{Name: upper.Bravo},
	)
	actualAlert.RawAlert = nil
	actualAlert.RawGroup = nil
	assert.Any(
		t,
		&rule.Rule{
			Name:  "Alfa",
			Group: "Bravo",
		},
		actualAlert,
	)
	actualRecord := rule.NewRecord(
		&v1.RecordingRule{Name: upper.Alfa},
		&v1.RuleGroup{Name: upper.Bravo},
	)
	actualRecord.RawRecord = nil
	actualRecord.RawGroup = nil
	assert.Any(
		t,
		&rule.Rule{
			Name:  "Alfa",
			Group: "Bravo",
		},
		actualRecord,
	)
}
