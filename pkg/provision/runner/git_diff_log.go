package runner

import "fmt"

func (r *Runner) gitDiffLog(
	old string,
	new string,
) string {
	c := r.newRun()
	c.Directory = r.clonePath
	c.Start(
		"git",
		"log",
		"--oneline",
		fmt.Sprintf("%s..%s", old, new),
		"--",
		r.toolPath,
	)

	return c.OutputString
}
