package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.GeneratedClient()

	for range 2 {
		r, e := c.GetListsWithResponse(context.Background(), nil)
		assert.FatalOnError(t, e)
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestRefusedTokenIsAnsweredAsForbiddenWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = forbidden.Format("linkace: %s", "Unauthenticated.")
	r, e := s.GeneratedClient().GetTagsWithResponse(context.Background(), nil)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRejectedLinkIsAnsweredAsBadRequestWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = validation.New(
		"linkace: %s",
		"The url field is required.",
	)
	r, e := s.GeneratedClient().AddLinkWithResponse(
		context.Background(),
		client.AddLinkJSONRequestBody{Link: "not a url"},
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON400)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestTakenTagNameIsAnsweredAsConflictWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = conflict.Format("linkace: %s", "The tag already exists.")
	r, e := s.GeneratedClient().CreateTagWithResponse(
		context.Background(),
		client.CreateTagJSONRequestBody{Name: "deploy"},
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON409)
	assert.Count(t, 0, s.Reporter.Events())
}
