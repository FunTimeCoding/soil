package peered

import "github.com/funtimecoding/soil/pkg/prometheus/face"

func (p *Peered) Alertmanager() face.AlertmanagerClient {
	return p.alertmanager
}
