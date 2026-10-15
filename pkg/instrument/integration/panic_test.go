package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/instrument/constant"
	"github.com/funtimecoding/soil/pkg/instrument/unit/sentry_receiver"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"os"
	"testing"
)

func TestEntrypointPanicIsReported(t *testing.T) {
	if os.Getenv(constant.FixtureChildEnvironment) != "" {
		i := instrument.New(
			identity.New(
				stringsConstant.LowerAlfa,
				stringsConstant.LowerAlfa,
				stringsConstant.LowerAlfa,
			),
		)
		defer func() { i.Flush(recover()) }()
		panic(stringsConstant.LowerAlfa)
	}

	received := sentry_receiver.New(t)
	r := run.New().NoPanic()
	r.Environment(constant.FixtureChildEnvironment, "1")
	r.Start(os.Args[0], "-test.run=^TestEntrypointPanicIsReported$")
	assert.Integer(t, 1, r.Exit)
	assert.Integer(t, 1, int(received.Load()))
}
