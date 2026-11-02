package daemon

import "context"

func (d *Daemon) Wait(x context.Context) error {
	if e := d.jiraGuard.Wait(x); e != nil {
		return e
	}

	return d.confluenceGuard.Wait(x)
}
