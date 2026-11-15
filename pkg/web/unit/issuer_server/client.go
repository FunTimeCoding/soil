package issuer_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (s *Server) Client(r *memory.Memory) *client.Client {
	return client.New(
		s.URL,
		"delta",
		"hotel",
		constant.SignInPath,
		"http://app.test/callback",
		client.DeriveKey("foxtrot"),
	).WithReporter(r)
}
