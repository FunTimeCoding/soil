package option

import "time"

func New() *Prometheus {
	return &Prometheus{Clock: time.Now, Client: BuildClient}
}
