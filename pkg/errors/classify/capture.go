package classify

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/caller_left"
	"github.com/funtimecoding/soil/pkg/face"
)

func Capture(
	x context.Context,
	r face.Reporter,
	e error,
	message string,
) (string, string) {
	text := Message(e, message)

	if caller_left.Is(e) {
		return text, ""
	}

	if c := Class(e); c != "" && !Reportable(e) {
		return text, r.Capture(x, fmt.Errorf("missing %s branch: %w", c, e))
	}

	return text, r.Capture(x, e)
}
