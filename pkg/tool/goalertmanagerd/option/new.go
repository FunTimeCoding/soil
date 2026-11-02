package option

func New() *Alertmanager {
	return &Alertmanager{
		Alertmanager: BuildAlertmanager,
		Prometheus:   BuildPrometheus,
	}
}
