package sentry_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	sentry1 "github.com/getsentry/sentry-go"
	"testing"
)

func RecordingHub(t *testing.T) (*sentry1.Hub, *sentry1.MockTransport) {
	t.Helper()
	transport := &sentry1.MockTransport{}
	client, e := sentry1.NewClient(
		sentry1.ClientOptions{
			Dsn:        "https://key@sentry.example/1",
			Transport:  transport,
			BeforeSend: sentry.Enrich,
		},
	)
	assert.FatalOnError(t, e)
	t.Cleanup(client.Close)

	return sentry1.NewHub(client, sentry1.NewScope()), transport
}
