package unit

import "github.com/funtimecoding/soil/pkg/failure"

func NewFixture() *Fixture {
	result := &Fixture{}
	result.failure = failure.New(result.Alfa)

	return result
}
