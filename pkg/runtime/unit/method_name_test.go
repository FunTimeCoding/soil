package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/runtime/unit/runtime_tester"
	"testing"
)

type named interface {
	Named() string
}

func TestMethodValueAndItsOwnCallerShareName(t *testing.T) {
	f := runtime_tester.NewMethodFixture()
	assert.String(t, runtime.MethodName(f.Named), f.Named())
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime/unit/runtime_tester.(*MethodFixture).Named",
		f.Named(),
	)
}

func TestMethodValueIsSameForEveryReceiver(t *testing.T) {
	assert.String(
		t,
		runtime.MethodName(runtime_tester.NewMethodFixture().Named),
		runtime.MethodName(runtime_tester.NewMethodFixture().Named),
	)
}

func TestMethodValueThroughInterfaceIsNotMethod(t *testing.T) {
	var f named = runtime_tester.NewMethodFixture()
	assert.True(
		t,
		runtime.MethodName(f.Named) != runtime.MethodName(runtime_tester.NewMethodFixture().Named),
	)
}
