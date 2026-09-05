package console_tester

import "github.com/funtimecoding/soil/pkg/tool/goclaude"

func (o *Tester) Check(session string) string {
	o.t.Helper()

	return goclaude.RunCheck(o.client, session)
}
