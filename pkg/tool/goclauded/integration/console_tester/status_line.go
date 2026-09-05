package console_tester

import "github.com/funtimecoding/soil/pkg/tool/goclaude"

func (o *Tester) StatusLine(body []byte) string {
	o.t.Helper()

	return goclaude.RunStatusLine(o.client, body)
}
