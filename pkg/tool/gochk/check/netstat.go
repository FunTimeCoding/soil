package check

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/linux/systemd/jc"
	"github.com/funtimecoding/soil/pkg/notation"
)

func Netstat(verbose bool) []*jc.Output {
	output := Execute(constant.NetstatCommand)

	if verbose {
		console.Format("Netstat raw: %s\n", output)
	}

	var result []*jc.Output
	notation.MustDecode(
		Pipe(
			Pipe(output, verbose, constant.Awk, "!seen[$4]++"),
			verbose,
			constant.Jc,
			"--netstat",
			"--monochrome",
		),
		&result,
		true,
	)

	return result
}
