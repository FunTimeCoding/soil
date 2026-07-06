package check

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linux"
	"github.com/funtimecoding/soil/pkg/linux/systemd/command"
	"github.com/funtimecoding/soil/pkg/linux/systemd/jc"
	"github.com/funtimecoding/soil/pkg/notation"
)

func Netstat(verbose bool) []*jc.Output {
	output := Execute(command.Netstat())

	if verbose {
		fmt.Printf("Netstat raw: %s\n", output)
	}

	var result []*jc.Output
	notation.MustDecode(
		Pipe(
			Pipe(output, verbose, linux.Awk, "!seen[$4]++"),
			verbose,
			linux.Jc,
			"--netstat",
			"--monochrome",
		),
		&result,
		true,
	)

	return result
}
