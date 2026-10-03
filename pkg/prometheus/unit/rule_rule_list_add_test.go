package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
	"time"
)

func TestRuleListAcceptsDuplicateRecordWithDifferingQuery(t *testing.T) {
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:  constant.UpperAlfa,
				Query: constant.LowerAlfa,
			},
			g,
		),
	)
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:  constant.UpperAlfa,
				Query: constant.LowerBravo,
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListAcceptsIdenticalRecordAcrossEvaluations(t *testing.T) {
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:           constant.UpperAlfa,
				Query:          constant.LowerAlfa,
				EvaluationTime: 1,
				LastEvaluation: time.Unix(1000, 0),
			},
			g,
		),
	)
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:           constant.UpperAlfa,
				Query:          constant.LowerAlfa,
				EvaluationTime: 2,
				LastEvaluation: time.Unix(2000, 0),
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListRejectsDriftedDuplicateRecord(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:   constant.UpperAlfa,
				Query:  constant.LowerAlfa,
				Health: v1.RuleHealthGood,
			},
			g,
		),
	)
	l.Add(
		rule.NewRecord(
			&v1.RecordingRule{
				Name:   constant.UpperAlfa,
				Query:  constant.LowerAlfa,
				Health: v1.RuleHealthBad,
			},
			g,
		),
	)
}

func TestRuleListAcceptsDuplicateAlertWithDifferingQuery(t *testing.T) {
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:  constant.UpperAlfa,
				Query: constant.LowerAlfa,
			},
			g,
		),
	)
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:  constant.UpperAlfa,
				Query: constant.LowerBravo,
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListAcceptsIdenticalAlertAcrossEvaluations(t *testing.T) {
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:           constant.UpperAlfa,
				Query:          constant.LowerAlfa,
				EvaluationTime: 1,
				LastEvaluation: time.Unix(1000, 0),
			},
			g,
		),
	)
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:           constant.UpperAlfa,
				Query:          constant.LowerAlfa,
				EvaluationTime: 2,
				LastEvaluation: time.Unix(2000, 0),
			},
			g,
		),
	)
	assert.Integer(t, 2, len(l.Get()))
}

func TestRuleListRejectsDriftedDuplicateAlert(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	g := &v1.RuleGroup{Name: constant.UpperBravo}
	l := rule_list.New()
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:   constant.UpperAlfa,
				Query:  constant.LowerAlfa,
				Health: v1.RuleHealthGood,
			},
			g,
		),
	)
	l.Add(
		rule.NewAlert(
			&v1.AlertingRule{
				Name:   constant.UpperAlfa,
				Query:  constant.LowerAlfa,
				Health: v1.RuleHealthBad,
			},
			g,
		),
	)
}
