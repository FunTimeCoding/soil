package argument_tester

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/unit/command_recorder"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/spf13/cobra"
	"testing"
)

func Instrumented(
	t *testing.T,
	root *cobra.Command,
	arguments ...string,
) *command_recorder.Recorder {
	t.Helper()
	result := command_recorder.New()
	argument.CobraInstrument(root, result)
	root.SetArgs(arguments)
	assert.FatalOnError(t, root.Execute())

	return result
}
