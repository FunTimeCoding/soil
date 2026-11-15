package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/unit/request_context_tester"
	"testing"
)

func TestAttributeOmitsBodyAndHeaders(t *testing.T) {
	m := request_context_tester.AttributeMap(
		request_context_tester.New().Attribute(),
	)
	assert.MapValue(t, "/hook", m, constant.TelemetryPath)
	assert.MapNotHasKey(t, m, constant.TelemetryBody)
	assert.MapNotHasKey(t, m, constant.TelemetryBodySize)
	assert.MapNotHasKey(t, m, "http.request.header.authorization")
}

func TestWebhookAttributeCarriesBodyAndHeaders(t *testing.T) {
	m := request_context_tester.AttributeMap(
		request_context_tester.New().WebhookAttribute(),
	)
	assert.MapValue(t, "/hook", m, constant.TelemetryPath)
	assert.MapValue(t, "payload", m, constant.TelemetryBody)
	assert.MapValue(t, int64(7), m, constant.TelemetryBodySize)
	assert.MapValue(
		t,
		[]string{"Bearer secret"},
		m,
		"http.request.header.authorization",
	)
}
