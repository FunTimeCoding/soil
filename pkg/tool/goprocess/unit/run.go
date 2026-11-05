package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/unit/command_end_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprocess"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"testing"
)

func run(
	t *testing.T,
	c *mock_client.Client,
	arguments ...string,
) (string, string, int) {
	t.Helper()
	var output bytes.Buffer
	var failure bytes.Buffer
	code := 0
	o := goprocess.New(
		func(string) face.Client { return c },
		terminal.NewWith(
			command_end_recorder.New(),
			&output,
			&failure,
			func(c int) { panic(exitSignal(c)) },
		),
	)
	o.SetArgs(arguments)
	func() {
		defer func() {
			if r := recover(); r != nil {
				c, okay := r.(exitSignal)

				if !okay {
					panic(r)
				}

				code = int(c)
			}
		}()
		assert.FatalOnError(t, o.Execute())
	}()

	return output.String(), failure.String(), code
}
