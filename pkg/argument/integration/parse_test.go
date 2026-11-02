package integration

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"os"
	"testing"
)

func TestAnUnknownFlagPrintsUsageAndExitsTwo(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		os.Args = split.Space(arguments)
		a := argument.NewInstance(
			identity.New("alfa", "alfa tool", "alfa [flags]"),
		)
		a.Boolean("bravo", false, "bravo flag")
		a.Parse()

		return
	}

	r := child(t, "alfa --charlie delta")
	assert.Integer(t, 2, r.Exit)
	assert.StringContains(t, "unknown flag: --charlie", r.ErrorString)
	assert.StringContains(t, "Usage:\n  alfa [flags]", r.ErrorString)
	assert.StringContains(t, "--bravo", r.ErrorString)
}

func TestASimpleInstanceRefusesAnUnknownFlagWithExitTwo(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		os.Args = split.Space(arguments)
		a := argument.NewSimple("alfa")
		a.Boolean("bravo", false, "bravo flag")
		a.ParseSimple()

		return
	}

	r := child(t, "alfa --charlie")
	assert.Integer(t, 2, r.Exit)
	assert.StringContains(t, "unknown flag: --charlie", r.ErrorString)
	assert.StringContains(t, "--bravo", r.ErrorString)
}

func TestASimpleInstanceAnswersHelpWithExitZero(t *testing.T) {
	if arguments := os.Getenv(constant.FixtureArgumentsEnvironment); arguments != "" {
		os.Args = split.Space(arguments)
		a := argument.NewSimple("alfa")
		a.Boolean("bravo", false, "bravo flag")
		a.ParseSimple()

		return
	}

	r := child(t, "alfa --help")
	assert.Integer(t, 0, r.Exit)
	assert.StringContains(t, "--bravo", r.ErrorString)
}
