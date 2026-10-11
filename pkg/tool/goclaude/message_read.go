package goclaude

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func messageRead(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "read <identifier>...",
		Short: "Read coordination messages whole by number",
		Args:  cobra.MinimumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			var identifiers []int

			for _, a := range arguments {
				identifiers = append(
					identifiers,
					int(strings.MustToUnsignedInteger(a)),
				)
			}

			text, e := ReadMessages(c.Client(), identifiers)
			errors.PanicOnError(e)
			console.Format("%s\n", text)
		},
	}
}
