package iterm_tester

import (
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"strings"
	"testing"
)

func New(
	t *testing.T,
	status int,
	answer string,
) *iterm.Client {
	t.Helper()
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(status),
		upstream_server.WithAnswer(answer),
	)
	result := iterm.New(iterm.WithHost(strings.TrimPrefix(s.URL, "http://")))
	t.Cleanup(result.Close)

	return result
}
