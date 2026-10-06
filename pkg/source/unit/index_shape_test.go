package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/reference_tester"
	"reflect"
	"testing"
)

func TestShapeFollowsFields(t *testing.T) {
	assert.True(
		t,
		index.Shape(reflect.TypeFor[reference_tester.Narrow]()) !=
			index.Shape(reflect.TypeFor[reference_tester.Wide]()),
	)
}

func TestShapeFollowsTags(t *testing.T) {
	assert.True(
		t,
		index.Shape(reflect.TypeFor[reference_tester.Narrow]()) !=
			index.Shape(reflect.TypeFor[reference_tester.Retagged]()),
	)
}

func TestShapeFollowsNestedFields(t *testing.T) {
	assert.True(
		t,
		index.Shape(reflect.TypeFor[reference_tester.HoldsNarrow]()) !=
			index.Shape(reflect.TypeFor[reference_tester.HoldsWide]()),
	)
}

func TestShapeIgnoresTypeName(t *testing.T) {
	assert.String(
		t,
		index.Shape(reflect.TypeFor[reference_tester.Narrow]()),
		index.Shape(reflect.TypeFor[reference_tester.Renamed]()),
	)
}
