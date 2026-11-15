package failure_tester

import "github.com/funtimecoding/soil/pkg/failure"

func New() *Tester {
	result := &Tester{}
	result.failure = failure.New(result.Alfa)

	return result
}
