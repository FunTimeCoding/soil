package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
	"time"
)

func TestAnUnknownChannelIsNotFound(t *testing.T) {
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/teams/tango/channels/name/alfa",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					w.WriteHeader(http.StatusNotFound)
					web.Encode(
						w,
						&model.AppError{
							Message:    "Unable to find the existing channel.",
							StatusCode: http.StatusNotFound,
						},
					)
				},
			)
		},
	)
	_, e := r.Client.GetChannelPosts(
		context.Background(),
		"alfa",
		&client.GetChannelPostsParams{Since: time.Time{}},
	)
	assert.True(t, not_found.Is(e))
}

func TestAnUpstreamFailureIsNotNotFound(t *testing.T) {
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/teams/tango/channels/name/alfa",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					w.WriteHeader(http.StatusUnauthorized)
					web.Encode(
						w,
						&model.AppError{
							Message:    "Invalid or expired session.",
							StatusCode: http.StatusUnauthorized,
						},
					)
				},
			)
		},
	)
	_, e := r.Client.GetChannelPosts(
		context.Background(),
		"alfa",
		&client.GetChannelPostsParams{Since: time.Time{}},
	)
	assert.True(t, e != nil)
	assert.False(t, not_found.Is(e))
	assert.StringContains(t, "Invalid or expired session.", e.Error())
}

func TestAnUnknownPostHasNoReactions(t *testing.T) {
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/posts/bravo",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					w.WriteHeader(http.StatusNotFound)
					web.Encode(
						w,
						&model.AppError{StatusCode: http.StatusNotFound},
					)
				},
			)
		},
	)
	_, e := r.Client.GetPostReactions(context.Background(), "bravo")
	assert.True(t, not_found.Is(e))
}

func TestTheClientReachesTheDaemon(t *testing.T) {
	r := client_tester.New(t, func(*http.ServeMux) {})
	assert.FatalOnError(t, r.Client.Reachable(context.Background()))
}
