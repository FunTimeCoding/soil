package argument_tester

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/unit/command_recorder"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/spf13/cobra"
)

func (o *Tester) Instrumented(
	root *cobra.Command,
	arguments ...string,
) *command_recorder.Recorder {
	o.t.Helper()
	result := command_recorder.New()
	argument.CobraInstrument(root, result)
	root.SetArgs(arguments)
	assert.FatalOnError(o.t, root.Execute())

	return result
}
