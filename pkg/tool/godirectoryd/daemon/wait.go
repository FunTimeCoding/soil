package daemon

import "context"

func (d *Daemon) Wait(x context.Context) error {
	return d.guard.Wait(x)
}
