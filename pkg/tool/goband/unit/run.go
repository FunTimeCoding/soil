package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/unit/command_end_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goband"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func run(
	t *testing.T,
	s *base.Server,
	arguments ...string,
) (string, string, int) {
	t.Helper()
	var output bytes.Buffer
	var failure bytes.Buffer
	code := 0
	o := goband.New(
		s.BandClient(t),
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
