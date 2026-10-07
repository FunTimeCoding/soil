package section

import "github.com/funtimecoding/soil/pkg/text/markdown/runbook/command"

type Section struct {
	Title    string
	Commands []*command.Command
}
