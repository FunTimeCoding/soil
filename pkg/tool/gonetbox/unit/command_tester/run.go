package command_tester

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/spf13/cobra"
	"testing"
)

func Run(
	t *testing.T,
	c *mock_client.Client,
	arguments ...string,
) (string, string, int) {
	t.Helper()

	return terminal_tester.Run(
		t,
		func(m *terminal.Terminal) *cobra.Command { return gonetbox.New(c, m) },
		arguments...,
	)
}
