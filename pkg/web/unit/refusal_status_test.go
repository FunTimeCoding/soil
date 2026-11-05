package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
	"testing"
)

func TestAnUpstreamsOwnRefusalsHaveNoClass(t *testing.T) {
	for _, status := range []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusUnprocessableEntity,
	} {
		assert.Nil(t, refusal.Status(status, "alfa"))
	}
}

func TestADaemonsUnprocessableStatusIsTheCallersMistake(t *testing.T) {
	e := refusal.NewDaemonTable().Classify(
		http.StatusUnprocessableEntity,
		"title is missing",
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "title is missing", e.Error())
}

func TestADaemonsForbiddenStatusIsForbidden(t *testing.T) {
	assert.True(
		t,
		forbidden.Is(
			refusal.NewDaemonTable().Classify(http.StatusForbidden, "not armed"),
		),
	)
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
