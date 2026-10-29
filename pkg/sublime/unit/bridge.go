package unit

import (
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"strings"
	"testing"
)

func bridge(
	t *testing.T,
	status int,
	answer string,
) *sublime.Client {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	result := sublime.New(
		sublime.WithHost(strings.TrimPrefix(s.URL, "http://")),
	)
	t.Cleanup(result.Close)

	return result
}
