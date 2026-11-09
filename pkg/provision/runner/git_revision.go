package runner

import "strings"

func (r *Runner) gitRevision(reference string) string {
	c := r.newRun()
	c.Directory = r.cloneDirectory
	c.Start("git", "rev-parse", reference)

	return strings.TrimSpace(c.OutputString)
}
