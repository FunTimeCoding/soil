package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/getsentry/sentry-go"
)

func Start(
	projectName string,
	locator string,
	version string,
) *sentry.Hub {
	errors.FatalOnEmpty(projectName, "projectName")
	errors.FatalOnEmpty(locator, "locator")

	if version == "" {
		version = constant.UndefinedVersion
	}

	h := sentry.CurrentHub()
	client, e := sentry.NewClient(
		sentry.ClientOptions{
			Dsn:              locator,
			Environment:      constant.UndefinedEnvironment,
			Release:          key_value.At(projectName, version),
			EnableTracing:    true,
			TracesSampleRate: 1.0,
			AttachStacktrace: true,
			SendDefaultPII:   true,
			BeforeSend:       enrichResponseBody,
		},
	)
	errors.FatalOnError(e)
	h.BindClient(client)

	return h
}
