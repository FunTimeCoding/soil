package client

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/opsgenie/opsgenie-go-sdk-v2/account"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/contact"
	"github.com/opsgenie/opsgenie-go-sdk-v2/deployment"
	"github.com/opsgenie/opsgenie-go-sdk-v2/escalation"
	"github.com/opsgenie/opsgenie-go-sdk-v2/forwarding_rule"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration"
	"github.com/opsgenie/opsgenie-go-sdk-v2/logs"
	"github.com/opsgenie/opsgenie-go-sdk-v2/maintenance"
	"github.com/opsgenie/opsgenie-go-sdk-v2/notification"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"github.com/opsgenie/opsgenie-go-sdk-v2/service"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
)

type Client struct {
	Account      *account.Client
	Alert        *alert.Client
	Contact      *contact.Client
	Deployment   *deployment.Client
	Escalation   *escalation.Client
	Forward      *forwarding_rule.Client
	Heartbeat    *heartbeat.Client
	Incident     *incident.Client
	Integration  *integration.Client
	Log          *logs.Client
	Maintenance  *maintenance.Client
	Notification *notification.Client
	Policy       *policy.Client
	Schedule     *schedule.Client
	Service      *service.Client
	Team         *team.Client
	User         *user.Client
	pool         *pool.Pool
}
