package integration

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"os"
	"testing"
)

func TestAnUnknownCobraFlagPrintsTheCommandUsageAndExitsTwo(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		argument.CobraExecute(cobraRoot(split.Space(arguments), nil))

		return
	}

	r := child(t, "bravo --delta")
	assert.Integer(t, 2, r.Exit)
	assert.StringContains(t, "unknown flag: --delta", r.ErrorString)
	assert.StringContains(t, "Usage:\n  alfa bravo", r.ErrorString)
	assert.StringContains(t, "--charlie", r.ErrorString)
	assert.StringNotContains(t, "panic", r.ErrorString)
}

func TestACobraCommandFailureStillPanics(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		argument.CobraExecute(
			cobraRoot(split.Space(arguments), errors.New("echo failed")),
		)

		return
	}

	r := child(t, "bravo")
	assert.StringContains(t, "panic: echo failed", r.ErrorString)
}
