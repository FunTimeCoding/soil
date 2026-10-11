package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/console_tester"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCheckCutsALongMessageThroughTheHook(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	receiver := c.Register("session-1")
	sender := c.Register("session-2")
	c.Check("session-1")
	s.Send(sender, receiver, strings.Repeat("a", 5000))
	output := c.Check("session-1")
	assert.StringContains(t, "…cut here - ", output)
	assert.StringContains(t, "read it with read_message", output)
	assert.True(t, utf8.RuneCountInString(output) <= 2000)
}

func TestReadMessagesThroughTheCommandLine(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	m := s.Store.SendMessage("Ash", "Cedar", "an answer")
	output := c.ReadMessages(int(m.Identifier), 999)
	assert.StringContains(
		t,
		fmt.Sprintf("[Message %d from Ash to Cedar at ", m.Identifier),
		output,
	)
	assert.StringContains(t, "]\nan answer\n", output)
	assert.StringContains(t, "No message with identifier: 999", output)
}
