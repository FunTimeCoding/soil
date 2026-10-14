//go:build unix

package reaper

import (
	"errors"
	"syscall"
)

func (r *Reaper) Interrupt() {
	r.owned.Range(
		func(
			key any,
			_ any,
		) bool {
			e := syscall.Kill(key.(int), syscall.SIGINT)

			if e != nil && !errors.Is(e, syscall.ESRCH) {
				r.reporter.CaptureException(e)
			}

			return true
		},
	)
}
