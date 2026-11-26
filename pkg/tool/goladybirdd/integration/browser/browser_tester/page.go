package browser_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/url"
)

func Page(m string) string {
	return join.Empty("data:text/html,", url.PathEscape(m))
}
