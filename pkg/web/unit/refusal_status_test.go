package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
	"testing"
)

func TestAnUnprocessableStatusIsTheCallersMistake(t *testing.T) {
	e := refusal.Status(http.StatusUnprocessableEntity, "title is missing")
	assert.True(t, validation.Is(e))
	assert.String(t, "title is missing", e.Error())
}

func TestARateLimitIsUnavailable(t *testing.T) {
	assert.True(
		t,
		unavailable.Is(refusal.Status(http.StatusTooManyRequests, "slow down")),
	)
}

func TestABadGatewayIsUnavailable(t *testing.T) {
	assert.True(
		t,
		unavailable.Is(refusal.Status(http.StatusBadGateway, "upstream gone")),
	)
}

func TestAGatewayTimeoutIsUnavailable(t *testing.T) {
	assert.True(
		t,
		unavailable.Is(
			refusal.Status(http.StatusGatewayTimeout, "upstream slow"),
		),
	)
}

func TestAnUnknownStatusHasNoClass(t *testing.T) {
	assert.Nil(t, refusal.Status(http.StatusTeapot, "short and stout"))
}
