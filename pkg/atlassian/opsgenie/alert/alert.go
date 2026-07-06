package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/detail"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user_map"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"time"
)

type Alert struct {
	Identifier        string
	SmallIdentifier   int
	Name              string
	Status            string
	Seen              bool
	Acknowledged      bool
	Snoozed           bool
	Priority          alert.Priority
	Link              string
	Owner             string
	Source            string
	Create            time.Time
	SnoozeUntil       time.Time
	Update            time.Time
	Report            alert.Report
	Responders        []alert.Responder
	Tags              []string
	Details           map[string]string
	Description       string
	Entity            string
	Category          string
	Team              *team.Team
	TeamKey           string
	TeamMap           *team_map.Map
	UserMap           *user_map.Map
	Prometheus        *detail.Prometheus
	shortUser         face.StringAlias
	shortAlert        face.StringAlias
	descriptionToName face.StringAlias
	parseDescription  constant.ParseDescription
	MonitorIdentifier string
	RawList           *alert.Alert
	RawDetail         *alert.GetAlertResult
}
