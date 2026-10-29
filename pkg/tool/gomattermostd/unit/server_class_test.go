package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	generated "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestAMissingPostIsAnsweredNotFound(t *testing.T) {
	r := memory.New()
	m := service_tester.NewMattermost(t, "kilo")
	s := server.New(service_tester.Reported(t, m, nil, nil, r), r)
	result, e := s.GetThread(
		context.Background(),
		generated.GetThreadRequestObject{Identifier: "alfa"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.GetThread404JSONResponse{Error: "post not found: alfa"},
		result,
	)
	assert.Count(t, 0, r.Events())
}

func TestARefusedPostIsAnsweredForbidden(t *testing.T) {
	r := memory.New()
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = forbidden.Format("no access to alfa")
	s := server.New(service_tester.Reported(t, m, nil, nil, r), r)
	result, e := s.GetPostReactions(
		context.Background(),
		generated.GetPostReactionsRequestObject{Identifier: "alfa"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.GetPostReactions403JSONResponse{Error: "no access to alfa"},
		result,
	)
	assert.Count(t, 0, r.Events())
}

func TestAnUnavailableMattermostIsAnsweredUnavailable(t *testing.T) {
	r := memory.New()
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = unavailable.Format("mattermost maintenance")
	s := server.New(service_tester.Reported(t, m, nil, nil, r), r)
	result, e := s.GetChannelPosts(
		context.Background(),
		generated.GetChannelPostsRequestObject{Name: "alfa"},
	)
	assert.FatalOnError(t, e)
	_, okay := result.(generated.GetChannelPosts503JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 0, r.Events())
}

func TestAMattermostFaultIsCaptured(t *testing.T) {
	r := memory.New()
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = fault()
	s := server.New(service_tester.Reported(t, m, nil, nil, r), r)
	result, e := s.GetThread(
		context.Background(),
		generated.GetThreadRequestObject{Identifier: "alfa"},
	)
	assert.FatalOnError(t, e)
	_, okay := result.(generated.GetThread500JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 1, r.Events())
}
