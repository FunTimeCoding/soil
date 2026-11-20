package command_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/spf13/cobra"
)

func (o *Tester) Run(
	c *mock_client.Client,
	arguments ...string,
) (string, string, int) {
	o.t.Helper()

	return terminal_tester.Run(
		o.t,
		func(m *terminal.Terminal) *cobra.Command {
			x := command_context.New(m, c, c)
			o.t.Cleanup(x.Close)

			return goclaude.New(x, memory.New())
		},
		arguments...,
	)
}
