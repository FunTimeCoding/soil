package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/result"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"testing"
)

func TestSilentCommandPrintsNothing(t *testing.T) {
	g := command_tester.New(t)

	for _, arguments := range command_tester.SilentCommands() {
		c := g.Mock()
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() {
				output, failure, code = g.Run(c, arguments...)
			},
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
	}
}

func TestFailedCommandExitsWithItsMessage(t *testing.T) {
	g := command_tester.New(t)

	for _, arguments := range append(
		command_tester.SilentCommands(),
		[]string{"run", "restart-all"},
		[]string{"run", "log", "alfa"},
		[]string{"run", soil.List},
		[]string{"run", "status"},
	) {
		c := g.Mock()
		c.Failure = errors.New("foxtrot")
		_, failure, code := g.Run(c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
	}
}

func TestStopReachesProcess(t *testing.T) {
	g := command_tester.New(t)
	c := g.Mock()
	_, _, code := g.Run(c, "run", "stop", "alfa")
	assert.Integer(t, 0, code)
	assert.False(t, c.ProcessList[0].Running)
}

func TestRestartAllEmitsCount(t *testing.T) {
	g := command_tester.New(t)
	output, _, code := g.Run(g.Mock(), "run", "restart-all")
	assert.Integer(t, 0, code)
	var r result.Restart
	notation.MustDecode(output, &r, false)
	assert.Integer(t, 2, r.Restarting)
}

func TestListEmitsNames(t *testing.T) {
	g := command_tester.New(t)
	output, _, code := g.Run(g.Mock(), "run", soil.List)
	assert.Integer(t, 0, code)
	var names []string
	notation.MustDecode(output, &names, false)
	assert.Strings(t, []string{"alfa", "bravo"}, names)
}

func TestStatusEmitsProcesses(t *testing.T) {
	g := command_tester.New(t)
	output, _, code := g.Run(g.Mock(), "run", "status")
	assert.Integer(t, 0, code)
	var processes []client.Process
	notation.MustDecode(output, &processes, false)
	assert.Count(t, 2, processes)
	assert.True(t, processes[0].Running)
	assert.False(t, processes[1].Running)
}

func TestLogPrintsLinesAndOlderHintApart(t *testing.T) {
	g := command_tester.New(t)
	c := g.Mock()
	c.Logs["alfa"] = []string{"kilo", "lima"}
	c.Older["alfa"] = 3
	output, failure, code := g.Run(c, "run", "log", "alfa")
	assert.Integer(t, 0, code)
	assert.String(t, "kilo\nlima\n", output)
	assert.String(t, "(3 older lines, use: log alfa all)\n", failure)
}

func TestLogAllAsksForEverythingAndPrintsNoHint(t *testing.T) {
	g := command_tester.New(t)
	c := g.Mock()
	c.Logs["alfa"] = []string{"kilo", "lima"}
	c.Older["alfa"] = 3
	output, failure, code := g.Run(c, "run", "log", "alfa", "all")
	assert.Integer(t, 0, code)
	assert.String(t, "kilo\nlima\n", output)
	assert.String(t, "", failure)
}

func TestUnknownProcessIsNotFound(t *testing.T) {
	g := command_tester.New(t)
	_, failure, code := g.Run(g.Mock(), "run", "restart", "charlie")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "charlie", failure)
}

func TestUnknownCommandIsRefused(t *testing.T) {
	g := command_tester.New(t)
	_, failure, code := g.Run(g.Mock(), "run", "wipe")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "unknown command wipe", failure)
}

func TestCheckEmitsSortedEntryNames(t *testing.T) {
	g := command_tester.New(t)
	path := join.Join(t.TempDir(), "Procfile")
	system.SaveFile(path, "bravo: sleep 1\nalfa: sleep 1\n")
	output, _, code := g.Run(g.Mock(), "--file", path, "check")
	assert.Integer(t, 0, code)
	var names []string
	notation.MustDecode(output, &names, false)
	assert.Strings(t, []string{"alfa", "bravo"}, names)
}
