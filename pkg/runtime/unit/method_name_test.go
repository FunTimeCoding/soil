package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"testing"
)

type named interface {
	Named() string
}

func TestMethodValueAndItsOwnCallerShareName(t *testing.T) {
	f := &MethodFixture{}
	assert.String(t, runtime.MethodName(f.Named), f.Named())
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime/unit.(*MethodFixture).Named",
		f.Named(),
	)
}

func TestMethodValueIsSameForEveryReceiver(t *testing.T) {
	assert.String(
		t,
		runtime.MethodName((&MethodFixture{}).Named),
		runtime.MethodName((&MethodFixture{}).Named),
	)
}

func TestMethodValueThroughInterfaceIsNotMethod(t *testing.T) {
	var f named = &MethodFixture{}
	assert.True(
		t,
		runtime.MethodName(f.Named) != runtime.MethodName((&MethodFixture{}).Named),
	)
}
