package client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

// New serves a mock Mattermost upstream, points the daemon's REST
// server at it, and returns a generated client against that server.
// The upstream pre-answers the team lookup for the "tango" team the
// client construction requires; the rest comes from configure.
func New(
	t *testing.T,
	configure func(*http.ServeMux),
) *Tester {
	t.Helper()
	upstream := mattermost_client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/teams/name/tango",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					web.Encode(w, &model.Team{Id: "tango", Name: "tango"})
				},
			)
			configure(m)
		},
		mattermost.WithTeam("tango"),
	)
	m := http.NewServeMux()
	generated.HandlerFromMux(
		generated.NewStrictHandler(
			server.New(upstream.Client, "test", memory.New()),
			nil,
		),
		m,
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	c, e := client.NewClientWithResponses(s.URL)
	errors.PanicOnError(e)

	return &Tester{Client: c}
}
