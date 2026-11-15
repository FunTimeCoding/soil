package runtime_tester

import "github.com/funtimecoding/soil/pkg/runtime"

func (f *MethodFixture) Named() string {
	return runtime.CallerName(0)
}
