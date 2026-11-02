package daemon

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Daemon) Close() {
	d.closeOnce.Do(
		func() {
			d.clock.Close()
			d.guard.Close()
			d.embedder.Close()
			errors.LogClose(d.reranker)
			errors.LogClose(d.counter)
			errors.PanicClose(d.database)
		},
	)
}
