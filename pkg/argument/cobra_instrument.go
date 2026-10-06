package argument

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/spf13/cobra"
)

func CobraInstrument(
	c *cobra.Command,
	r face.CommandRecorder,
) {
	before := c.PersistentPreRun
	c.PersistentPreRun = func(
		m *cobra.Command,
		arguments []string,
	) {
		r.BeginCommand(m.CommandPath())

		if before != nil {
			before(m, arguments)
		}
	}
	after := c.PersistentPostRun
	c.PersistentPostRun = func(
		m *cobra.Command,
		arguments []string,
	) {
		if after != nil {
			after(m, arguments)
		}

		r.RecordCommand(m.CommandPath())
	}
}
