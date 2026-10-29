package client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
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
	guard.New(m, []string{constant.ModelContextTestToken}).TokenMount(
		webConstant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(
					service_tester.Service(t, upstream.Client, nil, nil),
					memory.New(),
				),
				nil,
			),
			http.NewServeMux(),
		),
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	c := client.New(
		locator.New(host).Port(port).Insecure(),
		constant.ModelContextTestToken,
	)
	t.Cleanup(c.Close)

	return &Tester{t: t, Client: c}
}
