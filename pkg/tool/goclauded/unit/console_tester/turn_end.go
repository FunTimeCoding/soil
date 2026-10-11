package console_tester

import "github.com/funtimecoding/soil/pkg/tool/goclaude"

func (o *Tester) TurnEnd(session string) {
	o.t.Helper()
	goclaude.RunTurnEnd(o.client, session)
}
