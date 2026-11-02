package check

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/systemd/jc"
)

func Netstat(verbose bool) []*jc.Output {
	output := Execute(systemConstant.NetstatCommand)

	if verbose {
		console.Format("Netstat raw: %s\n", output)
	}

	var result []*jc.Output
	notation.MustDecode(
		Pipe(
			Pipe(output, verbose, systemConstant.Awk, "!seen[$4]++"),
			verbose,
			systemConstant.Jc,
			"--netstat",
			"--monochrome",
		),
		&result,
		true,
	)

	return result
}
