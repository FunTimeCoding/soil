package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestRestrictedIssueIsAnsweredForbidden(t *testing.T) {
	m, e := client_tester.New(t).GetIssue(
		forbidden.Format("You may not see ALFA"),
	)
	assert.True(t, forbidden.Is(e))
	assert.StringContains(t, "You may not see ALFA", e.Error())
	assert.Count(t, 0, m.Events())
}

func TestOutageIsAnsweredUnavailable(t *testing.T) {
	m, e := client_tester.New(t).GetIssue(outage_tester.RefusedDial())
	assert.True(t, unavailable.Is(e))
	assert.Count(t, 0, m.Events())
}

func TestFaultIsCaptured(t *testing.T) {
	m, e := client_tester.New(t).GetIssue(
		errors.Format("unexpected answer", "ALFA"),
	)
	assert.NotNil(t, e)
	assert.Count(t, 1, m.Events())
}
