package peered

func (p *Peered) Close() {
	p.alertmanagerGuard.Close()
	p.prometheusGuard.Close()
	p.alertmanager.Close()
	p.prometheus.Close()
}
