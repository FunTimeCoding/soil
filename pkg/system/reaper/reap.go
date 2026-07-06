package reaper

import (
	"fmt"
	"syscall"
)

func (r *Reaper) reap() {
	for pid, detail := range r.scan() {
		if r.Contains(pid) {
			continue
		}

		var status syscall.WaitStatus
		reaped, e := syscall.Wait4(pid, &status, syscall.WNOHANG, nil)

		if reaped <= 0 || e != nil {
			continue
		}

		r.reporter.CaptureWithContext(
			fmt.Errorf("zombie reaped"),
			"reaper",
			map[string]any{
				"pid":         reaped,
				"exit_status": status.ExitStatus(),
				"comm":        detail.comm,
				"ppid":        detail.ppid,
			},
		)
	}
}
