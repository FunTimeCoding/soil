package failure

import "github.com/funtimecoding/soil/pkg/runtime"

func New(method any) *Failure {
	return &Failure{
		owner:  owner(runtime.MethodName(method)),
		errors: map[string]error{},
	}
}
