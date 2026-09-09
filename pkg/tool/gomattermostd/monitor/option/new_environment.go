package option

import (
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"os"
)

func NewEnvironment() *Monitor {
	if !environment.Exists("MONITORING_ENABLED") {
		return &Monitor{Enabled: false}
	}

	return &Monitor{
		Enabled:  true,
		Schedule: environment.Required("MONITORING_SCHEDULE"),
		Channels: parseList(environment.Required("MONITORING_CHANNELS")),
		Topics:   parseList(environment.Required("MONITORING_TOPICS")),
		MessageLimit: strings.MustToInteger(
			environment.Required("MONITORING_MESSAGE_LIMIT"),
		),
		NotificationChannel: os.Getenv("MONITORING_NOTIFICATION_CHANNEL"),
		Username:            os.Getenv("MONITORING_USERNAME"),
	}
}
