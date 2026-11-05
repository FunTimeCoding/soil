package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/sublime/face"
	"github.com/funtimecoding/soil/pkg/sublime/mock_client"
	"reflect"
	"testing"
)

func TestTheSublimeClientAndItsMockShareTheFace(t *testing.T) {
	f := reflect.TypeFor[face.SublimeClient]()
	assert.True(t, reflect.TypeFor[*sublime.Client]().Implements(f))
	assert.True(t, reflect.TypeFor[*mock_client.Client]().Implements(f))
}
