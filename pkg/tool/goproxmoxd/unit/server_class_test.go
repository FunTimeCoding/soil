package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestARefusedPermissionAnswersForbidden(t *testing.T) {
	s := base.New(t)
	s.MockClient.SetFailure(
		forbidden.Format("Permission check failed (/, Sys.Audit)"),
	)
	r, e := restClient(t, s).ListNodesWithResponse(
		t.Context(),
		&client.ListNodesParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 403, r.StatusCode())
	assert.StringContains(t, "Permission check failed", r.JSON403.Error)
}

func TestALockedGuestAnswersConflict(t *testing.T) {
	s := base.New(t)
	s.MockClient.SetFailure(conflict.Format("VM is locked (backup)"))
	r, e := restClient(t, s).StartMachineWithResponse(
		t.Context(),
		101,
		&client.StartMachineParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 409, r.StatusCode())
}

func TestAnUnreachableHypervisorAnswersUnavailable(t *testing.T) {
	s := base.New(t)
	s.MockClient.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"test",
			"",
			errorsConstant.Refused,
		),
	)
	r, e := restClient(t, s).ListNodesWithResponse(
		t.Context(),
		&client.ListNodesParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 503, r.StatusCode())
}

func TestAnUnexpectedFailureIsCapturedWithItsEvent(t *testing.T) {
	s := base.New(t)
	s.MockClient.SetFailure(errors.New("cfs lock timeout"))
	r, e := restClient(t, s).ListNodesWithResponse(
		t.Context(),
		&client.ListNodesParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 500, r.StatusCode())
	assert.StringContains(t, "cfs lock timeout", r.JSON500.Error)
}
