package check

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
)

func Pipe(
	input string,
	verbose bool,
	s ...string,
) string {
	stdout, stderr := system.Pipe(input, s...)

	if verbose {
		if stdout != "" {
			console.Format("Pipe output: %s\n", stdout)
		}

		if stderr != "" {
			console.Format("Pipe error: %s\n", stderr)
		}
	}

	return stdout
}
