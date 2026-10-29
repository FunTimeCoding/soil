package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/connection"
)

func prometheusClient(b *board.Board) func() face.Client {
	return func() face.Client {
		return prometheus.New(
			b.Connection.Prometheus.Host,
			connection.Port(b.Connection.Prometheus),
			b.Connection.Prometheus.Secure,
			"",
			"",
		)
	}
}
