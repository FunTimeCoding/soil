package peered

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"

func (p *Peered) Alertmanager() face.Client {
	return p.alertmanager
}
