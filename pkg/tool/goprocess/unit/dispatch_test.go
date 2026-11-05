package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/result"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"testing"
)

func TestASilentCommandPrintsNothing(t *testing.T) {
	for _, arguments := range silentCommands() {
		c := newMock(t)
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() { output, failure, code = run(t, c, arguments...) },
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
	}
}

func TestAFailedCommandExitsWithItsMessage(t *testing.T) {
	for _, arguments := range append(
		silentCommands(),
		[]string{"run", "restart-all"},
		[]string{"run", "log", "alfa"},
		[]string{"run", soil.List},
		[]string{"run", "status"},
	) {
		c := newMock(t)
		c.Failure = errors.New("foxtrot")
		_, failure, code := run(t, c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
	}
}

func TestStopReachesTheProcess(t *testing.T) {
	c := newMock(t)
	_, _, code := run(t, c, "run", "stop", "alfa")
	assert.Integer(t, 0, code)
	assert.False(t, c.ProcessList[0].Running)
}

func TestRestartAllEmitsTheCount(t *testing.T) {
	output, _, code := run(t, newMock(t), "run", "restart-all")
	assert.Integer(t, 0, code)
	var r result.Restart
	notation.MustDecode(output, &r, false)
	assert.Integer(t, 2, r.Restarting)
}

func TestListEmitsTheNames(t *testing.T) {
	output, _, code := run(t, newMock(t), "run", soil.List)
	assert.Integer(t, 0, code)
	var names []string
	notation.MustDecode(output, &names, false)
	assert.Strings(t, []string{"alfa", "bravo"}, names)
}

func TestStatusEmitsTheProcesses(t *testing.T) {
	output, _, code := run(t, newMock(t), "run", "status")
	assert.Integer(t, 0, code)
	var processes []client.Process
	notation.MustDecode(output, &processes, false)
	assert.Count(t, 2, processes)
	assert.True(t, processes[0].Running)
	assert.False(t, processes[1].Running)
}

func TestLogPrintsTheLinesAndTheOlderHintApart(t *testing.T) {
	c := newMock(t)
	c.Logs["alfa"] = []string{"kilo", "lima"}
	c.Older["alfa"] = 3
	output, failure, code := run(t, c, "run", "log", "alfa")
	assert.Integer(t, 0, code)
	assert.String(t, "kilo\nlima\n", output)
	assert.String(t, "(3 older lines, use: log alfa all)\n", failure)
}

func TestLogAllAsksForEverythingAndPrintsNoHint(t *testing.T) {
	c := newMock(t)
	c.Logs["alfa"] = []string{"kilo", "lima"}
	c.Older["alfa"] = 3
	output, failure, code := run(t, c, "run", "log", "alfa", "all")
	assert.Integer(t, 0, code)
	assert.String(t, "kilo\nlima\n", output)
	assert.String(t, "", failure)
}

func TestAnUnknownProcessIsNotFound(t *testing.T) {
	_, failure, code := run(t, newMock(t), "run", "restart", "charlie")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "charlie", failure)
}

func TestAnUnknownCommandIsRefused(t *testing.T) {
	_, failure, code := run(t, newMock(t), "run", "wipe")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "unknown command wipe", failure)
}

func TestCheckEmitsTheSortedEntryNames(t *testing.T) {
	path := join.Join(t.TempDir(), "Procfile")
	system.SaveFile(path, "bravo: sleep 1\nalfa: sleep 1\n")
	output, _, code := run(t, newMock(t), "--file", path, "check")
	assert.Integer(t, 0, code)
	var names []string
	notation.MustDecode(output, &names, false)
	assert.Strings(t, []string{"alfa", "bravo"}, names)
}
