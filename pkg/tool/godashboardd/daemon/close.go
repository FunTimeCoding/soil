package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.prometheusGuard.Close()
			d.prometheus.Close()

			if d.argocd != nil {
				d.argocdGuard.Close()
				d.argocd.Close()
			}

			if d.usage != nil {
				d.usageGuard.Close()
				d.usage.Close()
			}

			d.authorization.Close()
			relational.CloseMapper(d.database)
		},
	)
}
