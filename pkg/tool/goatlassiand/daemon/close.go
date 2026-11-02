package daemon

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.jiraGuard.Close()
			d.confluenceGuard.Close()
			d.jira.Close()
			d.confluence.Close()
		},
	)
}
