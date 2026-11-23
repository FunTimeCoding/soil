package command_tester

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosalt"
	"github.com/spf13/cobra"
	"testing"
)

func Run(
	t *testing.T,
	arguments ...string,
) (string, string, int) {
	t.Helper()

	return terminal_tester.Run(
		t,
		func(m *terminal.Terminal) *cobra.Command {
			return gosalt.New(m)
		},
		arguments...,
	)
}
