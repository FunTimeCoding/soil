package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"testing"
)

func TestCapturedConnectionFailureGroupsByKindAndHost(t *testing.T) {
	h, transport := newRecordingHub(t)
	defer transport.Close()
	h.CaptureException(refused())
	event := transport.Events()[0]
	assert.Strings(
		t,
		[]string{"unreachable", "alfa.example:8080"},
		event.Fingerprint,
	)
	assert.String(t, "unreachable", event.Tags[constant.Kind])
	assert.String(t, "alfa.example:8080", event.Tags[constant.Host])
	assert.Any(
		t,
		"/api/items",
		event.Contexts[constant.Connection][constant.Path],
	)
	main := event.Exception[len(event.Exception)-1]
	assert.String(t, "unreachable", main.Type)
	assert.String(t, "alfa.example:8080: connection refused", main.Value)

	for _, x := range event.Exception {
		assert.StringNotContains(t, "secret", x.Value)
	}
}

func TestRecoveredConnectionFailureGroupsToo(t *testing.T) {
	h, transport := newRecordingHub(t)
	defer transport.Close()
	h.Recover(refused())
	event := transport.Events()[0]
	assert.Strings(
		t,
		[]string{"unreachable", "alfa.example:8080"},
		event.Fingerprint,
	)

	for _, x := range event.Exception {
		assert.StringNotContains(t, "secret", x.Value)
	}
}

func TestOtherErrorsKeepDefaultGrouping(t *testing.T) {
	h, transport := newRecordingHub(t)
	defer transport.Close()
	h.CaptureException(errorsNew("bad request"))
	assert.Count(t, 0, transport.Events()[0].Fingerprint)
}

func TestWrappedConnectionFailureLosesTheQuery(t *testing.T) {
	h, transport := newRecordingHub(t)
	defer transport.Close()
	h.CaptureException(fmt.Errorf("search: %w", refused()))
	event := transport.Events()[0]
	assert.True(t, len(event.Exception) > 1)

	for _, x := range event.Exception {
		assert.StringNotContains(t, "secret", x.Value)
	}
}
