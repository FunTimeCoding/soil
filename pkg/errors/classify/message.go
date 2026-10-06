package classify

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func Message(
	e error,
	message string,
) string {
	if not_found.Is(e) || unexpected.Is(e) || validation.Is(e) ||
		unreachable.Is(e) || timeout.Is(e) || dropped.Is(e) || expected(e) {
		return fmt.Sprintf("%s: %s", message, e)
	}

	return message
}
