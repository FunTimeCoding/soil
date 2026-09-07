package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"testing"
)

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
