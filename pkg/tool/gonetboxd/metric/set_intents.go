package metric

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/intent"

func (m *Metric) SetIntents(v []*intent.Intent) {
	m.status.Reset()

	for _, i := range v {
		m.status.WithLabelValues(i.Kind, i.Name, i.Status).Set(1)
	}
}
