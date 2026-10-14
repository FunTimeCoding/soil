package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"reflect"
	"testing"
)

func TestRunnerImplementsDrainer(t *testing.T) {
	assert.True(
		t,
		reflect.TypeFor[*runner.Runner]().Implements(
			reflect.TypeFor[face.Drainer](),
		),
	)
}
