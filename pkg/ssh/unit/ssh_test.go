package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/result"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "-T", constant.NoPTYArgument)
	assert.String(t, "-tt", constant.ForcePTYArgument)
	assert.String(t, "-v", constant.VerboseArgument)
}

func TestResult(t *testing.T) {
	assert.Any(
		t,
		&result.Result{
			OutputString: "Alfa",
			ErrorString:  "Bravo",
			Exit:         1,
			Error:        nil,
		},
		result.New(
			stringsConstant.UpperAlfa,
			stringsConstant.UpperBravo,
			1,
			nil,
		),
	)
}

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(
		t,
		"",
		ssh.EnvironmentPrefix(command.New(stringsConstant.UpperAlfa)),
	)
}

func TestPanicDefault(t *testing.T) {
	assert.True(t, ssh.NewWithPassword("alfa", "localhost", "", false).Panic)
	assert.False(
		t,
		ssh.NewWithPassword("alfa", "localhost", "", false).NoPanic().Panic,
	)
}

func TestNotDialed(t *testing.T) {
	assert.Panic(
		t,
		func() {
			ssh.NewWithPassword("alfa", "localhost", "", false).Run("true")
		},
	)
}

func TestDialUnreachable(t *testing.T) {
	e := ssh.NewWithPassword("alfa", "unknown.invalid", "", false).Dial()
	assert.True(t, unreachable.Is(e))
	assert.String(t, "unknown.invalid: unknown host", e.Error())
}
