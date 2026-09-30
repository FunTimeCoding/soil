package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/event"
	"github.com/funtimecoding/soil/pkg/face"
	"testing"
)

func TestEventErrorCarriesTheStory(t *testing.T) {
	wrapped := errors.New("index out of range")
	e := event.New("state_changed", `{"entity_id":"light.desk"}`, wrapped)
	assert.String(t, "event state_changed: index out of range", e.Error())
	assert.True(t, errors.Is(e, wrapped))
	assert.True(t, event.Is(e))
	var provider face.ContextProvider = e
	key, context := provider.ErrorContext()
	assert.String(t, "event", key)
	assert.String(t, "state_changed", context[constant.Type].(string))
	assert.String(
		t,
		`{"entity_id":"light.desk"}`,
		context[constant.Raw].(string),
	)
}
