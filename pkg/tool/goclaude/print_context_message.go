package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func printContextMessage(
	t *terminal.Terminal,
	m client.SessionMessage,
) {
	text, cut := cutRunes(m.Text, 300)

	if cut {
		text = fmt.Sprintf("%s…", text)
	}

	t.Line(fmt.Sprintf("**%s:** %s", m.Role, text))
	t.Line()
}
