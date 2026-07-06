package check

import (
	"fmt"
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
			fmt.Printf("Pipe output: %s\n", stdout)
		}

		if stderr != "" {
			fmt.Printf("Pipe error: %s\n", stderr)
		}
	}

	return stdout
}
