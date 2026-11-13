package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/shorten"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func printContextMessage(
	t *terminal.Terminal,
	m client.SessionMessage,
) {
	t.Line(fmt.Sprintf("**%s:** %s", m.Role, shorten.Ellipsis(m.Text, 300)))
	t.Line()
}
