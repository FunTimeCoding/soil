package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	m := memory.New()
	s := restClient(t, c, m)

	for range 2 {
		r, e := s.GetListsWithResponse(context.Background(), nil)
		assert.FatalOnError(t, e)
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, m.Events())
}

func TestRefusedTokenIsAnsweredAsForbiddenWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("linkace: %s", "Unauthenticated.")
	m := memory.New()
	r, e := restClient(t, c, m).GetTagsWithResponse(context.Background(), nil)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, m.Events())
}

func TestRejectedLinkIsAnsweredAsBadRequestWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("linkace: %s", "The url field is required.")
	m := memory.New()
	r, e := restClient(t, c, m).AddLinkWithResponse(
		context.Background(),
		client.AddLinkJSONRequestBody{Link: "not a url"},
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON400)
	assert.Count(t, 0, m.Events())
}

func TestTakenTagNameIsAnsweredAsConflictWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = conflict.Format("linkace: %s", "The tag already exists.")
	m := memory.New()
	r, e := restClient(t, c, m).CreateTagWithResponse(
		context.Background(),
		client.CreateTagJSONRequestBody{Name: "deploy"},
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON409)
	assert.Count(t, 0, m.Events())
}
