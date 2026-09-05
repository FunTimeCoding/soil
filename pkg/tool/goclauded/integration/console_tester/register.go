package console_tester

import "github.com/funtimecoding/soil/pkg/tool/goclaude"

func (o *Tester) Register(session string) string {
	o.t.Helper()

	return goclaude.RunRegister(o.client, session)
}
