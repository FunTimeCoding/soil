package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unreadable_body"
	"testing"
)

func TestUnreadableBodyCarriesTheCause(t *testing.T) {
	cause := errors.New("unexpected EOF")
	e := unreadable_body.New(cause, "technitium %s: %d", "/zones", 502)
	assert.String(
		t,
		"technitium /zones: 502 (body unreadable: unexpected EOF)",
		e.Error(),
	)
	assert.True(t, unreadable_body.Is(e))
	assert.True(t, errors.Is(e, cause))
	assert.False(t, unreadable_body.Is(cause))
}
