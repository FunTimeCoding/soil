package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/postgres"
	"github.com/funtimecoding/soil/pkg/postgres/constant"
	"github.com/funtimecoding/soil/pkg/postgres/refusal"
	"testing"
)

func TestSyntaxErrorIsCallersMistake(t *testing.T) {
	assert.True(t, validation.Is(refusal.Classify(refused("42601"))))
}

func TestDivisionByZeroIsCallersMistake(t *testing.T) {
	assert.True(t, validation.Is(refusal.Classify(refused("22012"))))
}

func TestMissingPrivilegeIsForbidden(t *testing.T) {
	assert.True(t, forbidden.Is(refusal.Classify(refused("42501"))))
}

func TestUniqueViolationIsConflict(t *testing.T) {
	assert.True(t, conflict.Is(refusal.Classify(refused("23505"))))
}

func TestSerializationFailureIsConflict(t *testing.T) {
	assert.True(t, conflict.Is(refusal.Classify(refused("40001"))))
}

func TestServerStartingUpIsUnavailable(t *testing.T) {
	assert.True(t, unavailable.Is(refusal.Classify(refused("57P03"))))
}

func TestTooManyConnectionsIsUnavailable(t *testing.T) {
	assert.True(t, unavailable.Is(refusal.Classify(refused("53300"))))
}

func TestCanceledQueryIsTimeout(t *testing.T) {
	assert.True(
		t,
		timeout.Is(refusal.Classify(refused(constant.QueryCanceledCode))),
	)
}

func TestBadPasswordIsOursToLookAt(t *testing.T) {
	assert.True(t, classify.Reportable(refusal.Classify(refused("28P01"))))
}

func TestInternalErrorIsOursToLookAt(t *testing.T) {
	assert.True(t, classify.Reportable(refusal.Classify(refused("XX000"))))
}

func TestCodeShorterThanClassIsOursToLookAt(t *testing.T) {
	assert.True(t, classify.Reportable(refusal.Classify(refused("4"))))
}

func TestErrorFromElsewherePassesThrough(t *testing.T) {
	e := errors.Format("elsewhere", "test")
	assert.True(t, refusal.Classify(e) == e)
}

func TestRefusedDialStaysConnectionFailure(t *testing.T) {
	c, e := postgres.New("postgres://alfa:bravo@127.0.0.1:1/charlie")
	assert.FatalOnError(t, e)
	t.Cleanup(c.Close)
	_, f := c.Query(t.Context(), "SELECT 1")
	assert.NotNil(t, connection.Classify(f))
}
