package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Client interface {
	Address() string
	Alerts(
		p *advanced_option.Alert,
		f *alert_filter.Filter,
		r *rule_list.List,
	) ([]*alert.Alert, *statistic.Statistic, error)
	Close()
	Create(
		name string,
		instance string,
		summary string,
		description string,
		generator string,
	) error
	DeleteSilence(identifier string) error
	Groups() (models.AlertGroups, error)
	MustAlerts(
		p *advanced_option.Alert,
		f *alert_filter.Filter,
		r *rule_list.List,
	) ([]*alert.Alert, *statistic.Statistic)
	MustCreate(
		name string,
		instance string,
		summary string,
		description string,
		generator string,
	)
	MustDeleteSilence(identifier string)
	MustSilences(expired bool) []*silence.Silence
	MustSimpleSilence(
		alert string,
		r *rule_list.List,
	) string
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
	SetSilence(
		alert string,
		comment string,
		d time.Duration,
		r *rule_list.List,
	) (string, error)
	SilenceByRule(name string) (*silence.Silence, error)
	SilenceExists(name string) (bool, error)
	Silences(expired bool) ([]*silence.Silence, error)
	SimpleSilence(
		alert string,
		r *rule_list.List,
	) (string, error)
	Status() (*models.AlertmanagerStatus, error)
}
