package integration

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/argument/unit/argument_tester"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"os"
	"testing"
)

func TestUnknownCobraFlagPrintsCommandUsageAndExitsTwo(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		argument.CobraExecute(argument_tester.Root(split.Space(arguments), nil))

		return
	}

	r := argument_tester.Child(t, "bravo --delta")
	assert.Integer(t, 2, r.Exit)
	assert.StringContains(t, "unknown flag: --delta", r.ErrorString)
	assert.StringContains(t, "Usage:\n  alfa bravo", r.ErrorString)
	assert.StringContains(t, "--charlie", r.ErrorString)
	assert.StringNotContains(t, "panic", r.ErrorString)
}

func TestCobraCommandFailureStillPanics(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		argument.CobraExecute(
			argument_tester.Root(
				split.Space(arguments),
				errors.New("echo failed"),
			),
		)

		return
	}

	r := argument_tester.Child(t, "bravo")
	assert.StringContains(t, "panic: echo failed", r.ErrorString)
}
