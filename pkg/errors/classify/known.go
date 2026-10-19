package classify

import "github.com/funtimecoding/soil/pkg/errors/unexpected"

func Known(e error) bool {
	return Answer(e) != nil || unexpected.Is(e) || expected(e)
}
