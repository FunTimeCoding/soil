package client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
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
			server.New(upstream.Client, memory.New()),
			nil,
		),
		m,
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	c, e := client.NewClientWithResponses(s.URL)
	errors.PanicOnError(e)

	return &Tester{t: t, Client: c}
}
