package sublime_tester

import (
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *sublime.Client {
	t.Helper()
	result := sublime.New(sublime.WithHost(web.TrimScheme(address)))
	t.Cleanup(result.Close)

	return result
}
