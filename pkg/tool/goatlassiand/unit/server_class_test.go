package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"testing"
)

func TestARestrictedIssueIsAnsweredForbidden(t *testing.T) {
	m, e := getIssue(t, forbidden.Format("You may not see ALFA"))
	assert.True(t, forbidden.Is(e))
	assert.StringContains(t, "You may not see ALFA", e.Error())
	assert.Count(t, 0, m.Events())
}

func TestAnOutageIsAnsweredUnavailable(t *testing.T) {
	m, e := getIssue(t, refusedConnection())
	assert.True(t, unavailable.Is(e))
	assert.Count(t, 0, m.Events())
}

func TestAFaultIsCaptured(t *testing.T) {
	m, e := getIssue(t, errors.Format("unexpected answer", "ALFA"))
	assert.NotNil(t, e)
	assert.Count(t, 1, m.Events())
}
