package option

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"os"
	"time"
)

func NewEnvironment() *Monitor {
	if !environment.Exists("MONITORING_ENABLED") {
		return &Monitor{Enabled: false}
	}

	interval, e := time.ParseDuration(
		environment.Required("MONITORING_INTERVAL"),
	)
	errors.PanicOnError(e)

	return &Monitor{
		Enabled:  true,
		Interval: interval,
		Channels: parseList(environment.Required("MONITORING_CHANNELS")),
		Topics:   parseList(environment.Required("MONITORING_TOPICS")),
		MessageLimit: strings.MustToInteger(
			environment.Required("MONITORING_MESSAGE_LIMIT"),
		),
		NotificationChannel: os.Getenv("MONITORING_NOTIFICATION_CHANNEL"),
		Username:            os.Getenv("MONITORING_USERNAME"),
	}
}
