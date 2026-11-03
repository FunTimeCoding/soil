//go:build unix

package reaper

import "syscall"

func (r *Reaper) Reap() {
	for pid, detail := range r.zombies() {
		if r.Contains(pid) {
			continue
		}

		var status syscall.WaitStatus
		reaped, e := syscall.Wait4(pid, &status, syscall.WNOHANG, nil)

		if reaped <= 0 || e != nil {
			continue
		}

		r.reaped.WithLabelValues(detail.Comm).Inc()
	}
}
