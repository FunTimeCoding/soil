package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
	"time"
)

func TestRuleListAcceptsDuplicateRecordWithDifferingQuery(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:  stringsConstant.UpperAlfa,
				Query: stringsConstant.LowerAlfa,
			},
			g,
		),
	)
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:  stringsConstant.UpperAlfa,
				Query: stringsConstant.LowerBravo,
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListAcceptsIdenticalRecordAcrossEvaluations(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:           stringsConstant.UpperAlfa,
				Query:          stringsConstant.LowerAlfa,
				EvaluationTime: 1,
				LastEvaluation: time.Unix(1000, 0),
			},
			g,
		),
	)
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:           stringsConstant.UpperAlfa,
				Query:          stringsConstant.LowerAlfa,
				EvaluationTime: 2,
				LastEvaluation: time.Unix(2000, 0),
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListRejectsDriftedDuplicateRecord(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:   stringsConstant.UpperAlfa,
				Query:  stringsConstant.LowerAlfa,
				Health: v1.RuleHealthGood,
			},
			g,
		),
	)
	assert.Panic(
		t,
		func() {
			l.Add(
				rule.NewRecord(
					&v1.RecordingRule{
						Name:   stringsConstant.UpperAlfa,
						Query:  stringsConstant.LowerAlfa,
						Health: v1.RuleHealthBad,
					},
					g,
				),
			)
		},
	)
}

func TestRuleListAcceptsDuplicateAlertWithDifferingQuery(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:  stringsConstant.UpperAlfa,
				Query: stringsConstant.LowerAlfa,
			},
			g,
		),
	)
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:  stringsConstant.UpperAlfa,
				Query: stringsConstant.LowerBravo,
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListAcceptsIdenticalAlertAcrossEvaluations(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:           stringsConstant.UpperAlfa,
				Query:          stringsConstant.LowerAlfa,
				EvaluationTime: 1,
				LastEvaluation: time.Unix(1000, 0),
			},
			g,
		),
	)
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:           stringsConstant.UpperAlfa,
				Query:          stringsConstant.LowerAlfa,
				EvaluationTime: 2,
				LastEvaluation: time.Unix(2000, 0),
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListRejectsDriftedDuplicateAlert(t *testing.T) {
	g := &v1.RuleGroup{Name: stringsConstant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:   stringsConstant.UpperAlfa,
				Query:  stringsConstant.LowerAlfa,
				Health: v1.RuleHealthGood,
			},
			g,
		),
	)
	assert.Panic(
		t,
		func() {
			l.Add(
				rule.NewAlert(
					&v1.AlertingRule{
						Name:   stringsConstant.UpperAlfa,
						Query:  stringsConstant.LowerAlfa,
						Health: v1.RuleHealthBad,
					},
					g,
				),
			)
		},
	)
}
