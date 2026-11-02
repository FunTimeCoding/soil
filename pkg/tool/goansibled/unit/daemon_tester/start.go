package daemon_tester

import "github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"

func Start(d *daemon.Daemon) {
	for _, w := range d.Workers() {
		w.Start()
	}
}
