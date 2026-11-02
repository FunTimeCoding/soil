package argument

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/spf13/cobra"
)

func CobraExecute(c *cobra.Command) {
	c.SetFlagErrorFunc(cobraFlagError)
	e := c.Execute()
	errors.PanicOnError(e)
}
