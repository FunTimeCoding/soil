package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestAnAccessDeniedFaultIsForbidden(t *testing.T) {
	c := band.New(
		faulting(t, http.StatusBadRequest, constant.FixtureAccessDeniedFault),
		constant.DefaultUser,
		"secret",
	)
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	assert.True(t, forbidden.Is(e))
}

func TestASenderFaultOfOurOwnEnvelopeIsOursToLookAt(t *testing.T) {
	c := band.New(
		faulting(
			t,
			http.StatusBadRequest,
			constant.FixtureInvalidParameterFault,
		),
		constant.DefaultUser,
		"secret",
	)
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	assert.True(t, classify.Reportable(e))
	assert.False(t, validation.Is(e))
}

func TestATimedOutFaultIsUnavailable(t *testing.T) {
	c := band.New(
		faulting(
			t,
			http.StatusInternalServerError,
			constant.FixtureTimedOutFault,
		),
		constant.DefaultUser,
		"secret",
	)
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	assert.True(t, unavailable.Is(e))
}

func TestAProxysServiceUnavailableIsUnavailable(t *testing.T) {
	c := band.New(
		faulting(t, http.StatusServiceUnavailable, ""),
		constant.DefaultUser,
		"secret",
	)
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	assert.True(t, unavailable.Is(e))
}

func TestARefusedCredentialIsOursToLookAt(t *testing.T) {
	c := band.New(
		faulting(t, http.StatusUnauthorized, ""),
		constant.DefaultUser,
		"secret",
	)
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	assert.True(t, classify.Reportable(e))
}

func TestARefusedDialIsAConnectionFailure(t *testing.T) {
	c := band.New("http://127.0.0.1:1", constant.DefaultUser, "secret")
	t.Cleanup(c.Close)
	_, e := c.PowerState()
	var f *connection.Failure
	assert.True(t, errors.As(e, &f))
}
