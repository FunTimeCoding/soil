package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
)

func TestRule(t *testing.T) {
	actualAlert := rule.NewAlert(
		&v1.AlertingRule{Name: stringsConstant.UpperAlfa},
		&v1.RuleGroup{Name: stringsConstant.UpperBravo},
	)
	actualAlert.RawAlert = nil
	actualAlert.RawGroup = nil
	assert.Any(t, &rule.Rule{Name: "Alfa", Group: "Bravo"}, actualAlert)
	actualRecord := rule.NewRecord(
		&v1.RecordingRule{Name: stringsConstant.UpperAlfa},
		&v1.RuleGroup{Name: stringsConstant.UpperBravo},
	)
	actualRecord.RawRecord = nil
	actualRecord.RawGroup = nil
	assert.Any(t, &rule.Rule{Name: "Alfa", Group: "Bravo"}, actualRecord)
}
