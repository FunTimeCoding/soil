package reaper

import (
	"fmt"
	"syscall"
)

func (r *Reaper) reap() {
	zombies := r.scan()

	for {
		var status syscall.WaitStatus
		pid, e := syscall.Wait4(-1, &status, syscall.WNOHANG, nil)

		if pid <= 0 || e != nil {
			return
		}

		context := map[string]any{
			"pid":         pid,
			"exit_status": status.ExitStatus(),
		}

		if detail, okay := zombies[pid]; okay {
			context["comm"] = detail.comm
			context["ppid"] = detail.ppid
		}

		r.reporter.CaptureWithContext(
			fmt.Errorf("zombie reaped"),
			"reaper",
			context,
		)
	}
}
