package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/spf13/cobra"
	"testing"
)

func runInstrumented(
	t *testing.T,
	root *cobra.Command,
	arguments ...string,
) *commandLog {
	t.Helper()
	result := &commandLog{}
	argument.CobraInstrument(root, result)
	root.SetArgs(arguments)
	assert.FatalOnError(t, root.Execute())

	return result
}
