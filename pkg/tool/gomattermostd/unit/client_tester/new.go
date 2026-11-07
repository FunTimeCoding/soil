package client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
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
	d := service_tester.NewDaemonWithoutSubscription(t, upstream.Client)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(c.Close)

	return &Tester{t: t, Client: c}
}
