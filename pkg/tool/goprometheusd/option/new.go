package option

func New() *Prometheus {
	return &Prometheus{Client: BuildClient}
}
