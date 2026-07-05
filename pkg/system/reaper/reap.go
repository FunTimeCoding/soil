package reaper

import (
	"fmt"
	"syscall"
)

func (r *Reaper) reap() {
	for {
		var status syscall.WaitStatus
		pid, e := syscall.Wait4(-1, &status, syscall.WNOHANG, nil)

		if pid <= 0 || e != nil {
			return
		}

		r.reporter.CaptureWithContext(
			fmt.Errorf("zombie reaped"),
			"reaper",
			map[string]any{
				"pid":         pid,
				"exit_status": status.ExitStatus(),
			},
		)
	}
}
