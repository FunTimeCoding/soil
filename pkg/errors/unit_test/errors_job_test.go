package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/job"
	"github.com/funtimecoding/soil/pkg/face"
	"testing"
)

func TestJobErrorCarriesTheStory(t *testing.T) {
	wrapped := errors.New("signal: killed")
	e := job.New(42, "palindrome", wrapped)
	e.Detail = map[string]any{"input": "clip.mp4"}
	assert.String(t, "job 42 (palindrome): signal: killed", e.Error())
	assert.True(t, errors.Is(e, wrapped))
	assert.True(t, job.Is(e))
	var provider face.ContextProvider = e
	key, context := provider.ErrorContext()
	assert.String(t, "job", key)
	assert.Integer(t, 42, context[constant.Identifier].(int))
	assert.String(t, "palindrome", context[constant.Kind].(string))
	assert.String(t, "clip.mp4", context["input"].(string))
}
