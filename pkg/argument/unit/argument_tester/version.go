package argument_tester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/spf13/cobra"
)

func (o *Tester) Version(arguments ...string) string {
	o.t.Helper()
	c := &cobra.Command{
		Use: "gotest",
		Run: func(
			*cobra.Command,
			[]string,
		) {
		},
	}
	argument.CobraStamp(c, identity.New("gotest", "test tool", "gotest"))
	var b bytes.Buffer
	c.SetOut(&b)
	c.SetArgs(arguments)
	assert.FatalOnError(o.t, c.Execute())

	return b.String()
}
