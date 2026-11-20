package web_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"maragu.dev/gomponents"
	"strings"
)

func (o *Tester) Render(n gomponents.Node) string {
	o.t.Helper()
	var b strings.Builder
	assert.FatalOnError(o.t, n.Render(&b))

	return b.String()
}
