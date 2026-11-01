package goagent

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/face"
	"github.com/spf13/cobra"
	"os"
)

func submit(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "submit <file>",
		Short: "Submit an intent file to the runner",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			content, e := os.ReadFile(arguments[0])
			errors.PanicOnError(e)
			t.EmitDone(c.PostRun(context.Background(), string(content)))
		},
	}
}
