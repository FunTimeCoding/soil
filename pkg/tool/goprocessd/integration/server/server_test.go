package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestStatusShowsRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send("status")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "*alfa", lines[0])
	assert.String(t, "*bravo", lines[1])
}

func TestListShowsAllProcesses(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send(constant.ListCommand)
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "alfa", lines[0])
	assert.String(t, "bravo", lines[1])
}

func TestRestartProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("restart", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}

func TestStopProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("stop", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "alfa", "status")
}

func TestStartStoppedProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.Send("stop", "alfa")
	s.WaitOutput(t, "alfa", "status")
	result := s.Send("start", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}

func TestRestartAllLeavesEveryProcessRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(
		t,
		"ok: restarting 2 processes in the background",
		s.Send("restart-all"),
	)
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	time.Sleep(time.Second)
	assert.String(t, "*alfa\n*bravo", s.Send("status"))
}

func TestRestartingEveryProcessOneByOneLeavesThemRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(t, "ok", s.Send("restart", "alfa"))
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(t, "ok", s.Send("restart", "bravo"))
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	time.Sleep(time.Second)
	assert.String(t, "*alfa\n*bravo", s.Send("status"))
}

func TestReloadProcfileAddsNewEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa\n*bravo", "status")
}

func TestReloadProcfileRemovesEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "alfa", constant.ListCommand)
}

func TestReloadProcfileChangedCommand(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 120\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}

func TestReloadProcfileUnchangedLeavesProcessRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, entry, "")
	s.WaitOutput(t, "*alfa", "status")
	assert.Integer(t, 1, launchCount(t, marker))
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitOutput(t, "*alfa", "status")
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}

func TestReloadProcfileRemovingOneLeavesOthersRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, fmt.Sprintf("%sbravo: sleep 60\n", entry), "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	s.WriteProcfile(entry)
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitOutput(t, "*alfa", "status")
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}

func TestReloadEnvironment(t *testing.T) {
	s := tester.New(
		t,
		"alfa: sh -c 'echo $TEST_VALUE && sleep 60'\n",
		"export TEST_VALUE=original\n",
	)
	s.WaitContains(t, "original", "log", "alfa")
	s.WriteEnvrc("export TEST_VALUE=updated\n")
	s.Send("reload-environment")
	s.Send("restart", "alfa")
	s.WaitContains(t, "updated", "log", "alfa")
}

func TestReloadEnvironmentRemovesVanishedExport(t *testing.T) {
	t.Setenv("TEST_REMOVED", "ghost")
	s := tester.New(
		t,
		"alfa: sh -c 'echo removed=${TEST_REMOVED:-gone} && sleep 60'\n",
		"export TEST_REMOVED=ghost\n",
	)
	s.WaitContains(t, "removed=ghost", "log", "alfa")
	s.WriteEnvrc("export TEST_OTHER=1\n")
	s.Send("reload-environment")
	s.Send("restart", "alfa")
	s.WaitContains(t, "removed=gone", "log", "alfa")
}

func TestLogCapturesOutput(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo hello-from-alfa && sleep 60'\n", "")
	s.WaitContains(t, "hello-from-alfa", "log", "alfa")
}

func TestLogUnknownProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("log", "nonexistent")
	assert.True(t, strings.HasPrefix(result, "error:"))
}

func TestLogShowsCurrentGenerationAfterRestart(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.Send("restart", "alfa")
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	result := s.Send("log", "alfa")
	assert.Integer(t, 1, strings.Count(result, "mark"))
	assert.True(t, strings.Contains(result, "(2 older lines"))
	assert.StringNotContains(t, "Terminating", result)
}

func TestLogAllShowsEveryGeneration(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.Send("restart", "alfa")
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	result := s.Send("log", "alfa", "all")
	assert.Integer(t, 2, strings.Count(result, "mark"))
	assert.True(t, strings.Contains(result, "Terminating alfa"))
}

func TestLogClearDiscardsHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	assert.String(t, "ok", s.Send("log", "alfa", "clear"))
	assert.String(t, "ok", s.Send("log", "alfa", "all"))
}

func TestLogUnknownOption(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.True(t, strings.HasPrefix(s.Send("log", "alfa", "wipe"), "error:"))
}

func TestReloadProcfileKeepsLogHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.WriteProcfile("alfa: sh -c 'echo mark && sleep 120'\n")
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	assert.Integer(t, 2, strings.Count(s.Send("log", "alfa", "all"), "mark"))
}

func TestProcessFailToStartDoesNotCrashManager(t *testing.T) {
	s := tester.New(t, "broken: /nonexistent/binary\nhealthy: sleep 60\n", "")
	s.WaitOutput(t, "broken\n*healthy", "status")
}

func TestProcessCrashMidSessionDoesNotCrashManager(t *testing.T) {
	s := tester.New(t, "short: sleep 0.1\nhealthy: sleep 60\n", "")
	s.WaitOutput(t, "short\n*healthy", "status")
}

func TestStatusDuringReloadProcfile(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	stop := make(chan struct{})
	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()

		for {
			select {
			case <-stop:
				return
			default:
				s.Send("status")
			}
		}
	}()

	for i := range 20 {
		if i%2 == 0 {
			s.WriteProcfile(
				"alfa: sleep 60\nbravo: sleep 60\ncharlie: sleep 60\n",
			)
		} else {
			s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
		}

		s.Send("reload-procfile")
	}

	close(stop)
	group.Wait()
}
