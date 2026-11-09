package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestBusyEngineAnswersUnavailable(t *testing.T) {
	s := base.New(t)
	s.Alfa.Fail(unavailable.Format("band power: fault the service is busy"))
	r, e := s.Client.GetStatusWithResponse(
		t.Context(),
		&client.GetStatusParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 503, r.StatusCode())
	assert.StringContains(t, "the service is busy", r.JSON503.Error)
}

func TestEngineRefusingAccountAnswersForbidden(t *testing.T) {
	s := base.New(t)
	s.Alfa.Fail(forbidden.Format("band power: fault access denied"))
	r, e := s.Client.RequestPowerWithResponse(
		t.Context(),
		client.RequestPowerJSONRequestBody{
			Instance:  new("alfa"),
			Operation: client.PowerRequestOperationShutdown,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 403, r.StatusCode())
	assert.StringContains(t, "access denied", r.JSON403.Error)
}
