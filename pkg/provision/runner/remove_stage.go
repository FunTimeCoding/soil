package runner

import "os"

func (r *Runner) removeStage(path string) {
	if e := os.RemoveAll(path); e != nil {
		r.reporter.CaptureException(e)
	}
}
