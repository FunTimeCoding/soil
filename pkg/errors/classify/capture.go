package classify

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
)

func Capture(
	r face.Reporter,
	e error,
	message string,
) (string, string) {
	text := Message(e, message)

	if c := Class(e); c != "" && !Reportable(e) {
		return text, r.CaptureException(
			fmt.Errorf("missing %s branch: %w", c, e),
		)
	}

	return text, r.CaptureException(e)
}
