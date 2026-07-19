package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
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
