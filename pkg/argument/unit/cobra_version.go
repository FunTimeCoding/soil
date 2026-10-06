package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/spf13/cobra"
	"testing"
)

func cobraVersion(
	t *testing.T,
	arguments ...string,
) string {
	t.Helper()
	o := &cobra.Command{Use: "gotest", Run: func(*cobra.Command, []string) {}}
	argument.CobraStamp(o, identity.New("gotest", "test tool", "gotest"))
	var b bytes.Buffer
	o.SetOut(&b)
	o.SetArgs(arguments)
	assert.FatalOnError(t, o.Execute())

	return b.String()
}
