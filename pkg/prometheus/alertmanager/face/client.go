package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Client interface {
	Address() string
	Alerts(
		p *advanced_option.Alert,
		f *alert_filter.Filter,
	) ([]*alert.Alert, *statistic.Statistic, error)
	Close()
	Create(
		name string,
		instance string,
		summary string,
		description string,
		expression string,
	) error
	DeleteSilence(identifier string) error
	Groups() (models.AlertGroups, error)
	MustAlerts(
		p *advanced_option.Alert,
		f *alert_filter.Filter,
	) ([]*alert.Alert, *statistic.Statistic)
	MustCreate(
		name string,
		instance string,
		summary string,
		description string,
		expression string,
	)
	MustDeleteSilence(identifier string)
	MustRule(name string) *rule.Rule
	MustRules() *rule_list.List
	MustSilences(expired bool) []*silence.Silence
	MustSimpleSilence(alert string) string
	MustStatus() *models.AlertmanagerStatus
	PostSilence(
		identifier string,
		alert string,
		comment string,
		start time.Time,
		end time.Time,
	) (string, error)
	Reachable(x context.Context) error
	Receivers() ([]*models.Receiver, error)
	Rule(name string) (*rule.Rule, error)
	RuleExists(name string) (bool, error)
	Rules() (*rule_list.List, error)
	Runbook(name string) (string, error)
	SetSilence(
		alert string,
		comment string,
		d time.Duration,
	) (string, error)
	SilenceByRule(name string) (*silence.Silence, error)
	SilenceExists(name string) (bool, error)
	Silences(expired bool) ([]*silence.Silence, error)
	SimpleSilence(alert string) (string, error)
	Status() (*models.AlertmanagerStatus, error)
}
