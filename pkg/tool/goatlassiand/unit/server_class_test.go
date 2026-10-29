package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"testing"
)

func TestARestrictedIssueIsAnsweredForbidden(t *testing.T) {
	r, m := getIssue(t, forbidden.Format("You may not see ALFA"))
	assert.Any(
		t,
		server.GetIssue403JSONResponse{Error: "You may not see ALFA"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestAnOutageIsAnsweredUnavailable(t *testing.T) {
	r, m := getIssue(t, refusedConnection())
	_, okay := r.(server.GetIssue503JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 0, m.Events())
}

func TestAFaultIsCaptured(t *testing.T) {
	r, m := getIssue(t, errors.Format("unexpected answer", "ALFA"))
	_, okay := r.(server.GetIssue500JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 1, m.Events())
}
