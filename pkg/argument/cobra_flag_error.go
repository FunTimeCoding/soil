package argument

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/spf13/cobra"
	"os"
)

func cobraFlagError(
	c *cobra.Command,
	e error,
) error {
	writer.Print(os.Stderr, "%s\n\n%s", e, c.UsageString())
	os.Exit(2)

	return e
}
