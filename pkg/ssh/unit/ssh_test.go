package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	sshConstant "github.com/funtimecoding/soil/pkg/ssh/constant"
	"github.com/funtimecoding/soil/pkg/ssh/result"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "-T", sshConstant.NoPTYArgument)
	assert.String(t, "-tt", sshConstant.ForcePTYArgument)
	assert.String(t, "-v", sshConstant.VerboseArgument)
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
		result.New(constant.UpperAlfa, constant.UpperBravo, 1, nil),
	)
}

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(t, "", ssh.EnvironmentPrefix(command.New(constant.UpperAlfa)))
}
