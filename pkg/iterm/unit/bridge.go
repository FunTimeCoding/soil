package unit

import (
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"strings"
	"testing"
)

func bridge(
	t *testing.T,
	status int,
	answer string,
) *iterm.Client {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	result := iterm.New(iterm.WithHost(strings.TrimPrefix(s.URL, "http://")))
	t.Cleanup(result.Close)

	return result
}
