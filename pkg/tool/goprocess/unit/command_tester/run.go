package command_tester

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/goprocess"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
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
			return goprocess.New(func(string) face.Client { return c }, m)
		},
		arguments...,
	)
}
