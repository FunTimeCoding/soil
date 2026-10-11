package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestReadMessageTool(t *testing.T) {
	s := base.New(t)
	direct := s.Store.SendMessage("Ash", "Cedar", "an answer")
	broadcast := s.Store.SendMessage("Dale", "", "deploying now")
	a := s.NewSession(t)
	a.Announce(a.Name(), "reading")
	result := a.MustCallTool(
		constant.ReadMessage,
		map[string]any{
			constant.Identifiers: []any{direct.Identifier, broadcast.Identifier},
		},
	)
	assert.StringContains(
		t,
		fmt.Sprintf("[Message %d from Ash to Cedar at ", direct.Identifier),
		result,
	)
	assert.StringContains(t, "]\nan answer\n", result)
	assert.StringContains(
		t,
		fmt.Sprintf(
			"[Message %d from Dale to everyone at ",
			broadcast.Identifier,
		),
		result,
	)
	assert.StringContains(t, "]\ndeploying now\n", result)
}

func TestReadMessageToolNamesUnknownNumbers(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "reading")
	result := a.MustCallTool(
		constant.ReadMessage,
		map[string]any{constant.Identifiers: []any{404, 405}},
	)
	assert.StringContains(t, "No message with identifier: 404,405", result)
}

func TestReadMessageToolRequiresIdentifiers(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "reading")
	message, e := a.CallToolError(
		constant.ReadMessage,
		map[string]any{constant.Identifiers: []any{}},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "identifiers is required", message)
}
