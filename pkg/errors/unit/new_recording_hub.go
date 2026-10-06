package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/getsentry/sentry-go"
	"testing"
)

func newRecordingHub(t *testing.T) (*sentry.Hub, *sentry.MockTransport) {
	t.Helper()
	transport := &sentry.MockTransport{}
	client, e := sentry.NewClient(
		sentry.ClientOptions{
			Dsn:        "https://key@sentry.example/1",
			Transport:  transport,
			BeforeSend: soil.Enrich,
		},
	)
	assert.FatalOnError(t, e)

	return sentry.NewHub(client, sentry.NewScope()), transport
}
