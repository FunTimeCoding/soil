package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestStatusShowsRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.Wait(t, "alfa,bravo", s.Running)
}

func TestListShowsAllProcesses(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	assert.String(t, "alfa,bravo", s.Names())
}

func TestTheClientReachesTheSocket(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.FatalOnError(t, s.Client.Reachable(context.Background()))
}

func TestRestartProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.Wait(t, "alfa", s.Running)
}

func TestStopProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.FatalOnError(t, s.Client.Stop(context.Background(), "alfa"))
	s.Wait(t, "", s.Running)
}

func TestStartStoppedProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.FatalOnError(t, s.Client.Stop(context.Background(), "alfa"))
	s.Wait(t, "", s.Running)
	assert.FatalOnError(t, s.Client.Start(context.Background(), "alfa"))
	s.Wait(t, "alfa", s.Running)
}

func TestAnUnknownProcessIsNotFound(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	x := context.Background()
	assert.True(t, not_found.Is(s.Client.Restart(x, "nonexistent")))
	assert.True(t, not_found.Is(s.Client.Stop(x, "nonexistent")))
	assert.True(t, not_found.Is(s.Client.Start(x, "nonexistent")))
	assert.True(t, not_found.Is(s.Client.ClearLog(x, "nonexistent")))
	_, e := s.Client.Log(x, "nonexistent", false)
	assert.True(t, not_found.Is(e))
}

func TestRestartAllLeavesEveryProcessRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.Wait(t, "alfa,bravo", s.Running)
	count, e := s.Client.RestartAll(context.Background())
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, count)
	s.Wait(t, "alfa,bravo", s.Running)
	time.Sleep(time.Second)
	assert.String(t, "alfa,bravo", s.Running())
}

func TestASecondWaveIsAConflict(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.Wait(t, "alfa,bravo", s.Running)
	_, e := s.Client.RestartAll(context.Background())
	assert.FatalOnError(t, e)
	_, f := s.Client.RestartAll(context.Background())
	assert.True(t, conflict.Is(f))
	s.Wait(t, "alfa,bravo", s.Running)
}

func TestRestartingEveryProcessOneByOneLeavesThemRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.Wait(t, "alfa,bravo", s.Running)
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.Wait(t, "alfa,bravo", s.Running)
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "bravo"))
	s.Wait(t, "alfa,bravo", s.Running)
	time.Sleep(time.Second)
	assert.String(t, "alfa,bravo", s.Running())
}

func TestReloadProcfileAddsNewEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.Wait(t, "alfa,bravo", s.Running)
}

func TestReloadProcfileRemovesEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\n")
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.Wait(t, "alfa", s.Names)
}

func TestReloadProcfileChangedCommand(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 120\n")
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.Wait(t, "alfa", s.Running)
}

func TestReloadProcfileUnchangedLeavesProcessRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, entry, "")
	s.Wait(t, "alfa", s.Running)
	assert.Integer(t, 1, launchCount(t, marker))
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.Wait(t, "alfa", s.Running)
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}

func TestReloadProcfileRemovingOneLeavesOthersRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, fmt.Sprintf("%sbravo: sleep 60\n", entry), "")
	s.Wait(t, "alfa,bravo", s.Running)
	s.WriteProcfile(entry)
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.Wait(t, "alfa", s.Running)
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}

func TestReloadEnvironment(t *testing.T) {
	s := tester.New(
		t,
		"alfa: sh -c 'echo $TEST_VALUE && sleep 60'\n",
		"export TEST_VALUE=original\n",
	)
	s.WaitLog(t, "original", "alfa")
	s.WriteEnvrc("export TEST_VALUE=updated\n")
	assert.FatalOnError(t, s.Client.ReloadEnvironment(context.Background()))
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.WaitLog(t, "updated", "alfa")
}

func TestReloadEnvironmentRemovesVanishedExport(t *testing.T) {
	t.Setenv("TEST_REMOVED", "ghost")
	s := tester.New(
		t,
		"alfa: sh -c 'echo removed=${TEST_REMOVED:-gone} && sleep 60'\n",
		"export TEST_REMOVED=ghost\n",
	)
	s.WaitLog(t, "removed=ghost", "alfa")
	s.WriteEnvrc("export TEST_OTHER=1\n")
	assert.FatalOnError(t, s.Client.ReloadEnvironment(context.Background()))
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.WaitLog(t, "removed=gone", "alfa")
}

func TestLogCapturesOutput(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo hello-from-alfa && sleep 60'\n", "")
	s.WaitLog(t, "hello-from-alfa", "alfa")
}

func TestLogShowsCurrentGenerationAfterRestart(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitLog(t, "mark", "alfa")
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.WaitLog(t, "mark", "alfa")
	l := s.Log("alfa", false)
	text := join.NewLine(l.Lines)
	assert.Integer(t, 1, strings.Count(text, "mark"))
	assert.Integer(t, 2, l.Older)
	assert.StringNotContains(t, "Terminating", text)
}

func TestLogAllShowsEveryGeneration(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitLog(t, "mark", "alfa")
	assert.FatalOnError(t, s.Client.Restart(context.Background(), "alfa"))
	s.WaitLog(t, "mark", "alfa")
	text := join.NewLine(s.Log("alfa", true).Lines)
	assert.Integer(t, 2, strings.Count(text, "mark"))
	assert.True(t, strings.Contains(text, "Terminating alfa"))
}

func TestLogClearDiscardsHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitLog(t, "mark", "alfa")
	assert.FatalOnError(t, s.Client.ClearLog(context.Background(), "alfa"))
	assert.Count(t, 0, s.Log("alfa", true).Lines)
}

func TestReloadProcfileKeepsLogHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitLog(t, "mark", "alfa")
	s.WriteProcfile("alfa: sh -c 'echo mark && sleep 120'\n")
	assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	s.WaitLog(t, "mark", "alfa")
	s.Wait(
		t,
		"2",
		func() string {
			return fmt.Sprint(
				strings.Count(join.NewLine(s.Log("alfa", true).Lines), "mark"),
			)
		},
	)
}

func TestProcessFailToStartDoesNotCrashManager(t *testing.T) {
	s := tester.New(t, "broken: /nonexistent/binary\nhealthy: sleep 60\n", "")
	s.Wait(t, "healthy", s.Running)
}

func TestProcessCrashMidSessionDoesNotCrashManager(t *testing.T) {
	s := tester.New(t, "short: sleep 0.1\nhealthy: sleep 60\n", "")
	s.Wait(t, "healthy", s.Running)
}

func TestStatusDuringReloadProcfile(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.Wait(t, "alfa,bravo", s.Running)
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
				s.Running()
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

		assert.FatalOnError(t, s.Client.ReloadProcfile(context.Background()))
	}

	close(stop)
	group.Wait()
}
