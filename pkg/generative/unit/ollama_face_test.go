package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/mock_client"
	"reflect"
	"testing"
)

func TestOllamaClientImplementsTheFace(t *testing.T) {
	assert.True(
		t,
		reflect.TypeFor[*ollama.Client]().Implements(
			reflect.TypeFor[face.OllamaClient](),
		),
	)
}

func TestOllamaMockImplementsTheFace(t *testing.T) {
	assert.True(
		t,
		reflect.TypeFor[*mock_client.Client]().Implements(
			reflect.TypeFor[face.OllamaClient](),
		),
	)
}
