package console_tester

import "github.com/funtimecoding/soil/pkg/tool/goclaude"

func (o *Tester) Statusline(body []byte) string {
	o.t.Helper()

	return goclaude.RunStatusline(o.client, body)
}
