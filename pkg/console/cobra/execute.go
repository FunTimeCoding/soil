package cobra

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/spf13/cobra"
)

func Execute(c *cobra.Command) {
	if e := c.Execute(); e != nil {
		system.Exitf(1, "command failed: %s\n", e)
	}
}
