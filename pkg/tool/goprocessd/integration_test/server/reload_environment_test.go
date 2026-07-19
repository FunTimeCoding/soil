package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestReloadEnvironment(t *testing.T) {
	s := tester.New(
		t,
		"alfa: sh -c 'echo $TEST_VALUE && sleep 60'\n",
		"export TEST_VALUE=original\n",
	)
	time.Sleep(200 * time.Millisecond)
	log := s.Send("log", "alfa")
	assert.True(t, strings.Contains(log, "original"))
	s.WriteEnvrc("export TEST_VALUE=updated\n")
	s.Send("reload-environment")
	s.Send("restart", "alfa")
	time.Sleep(200 * time.Millisecond)
	log = s.Send("log", "alfa")
	assert.True(t, strings.Contains(log, "updated"))
}
