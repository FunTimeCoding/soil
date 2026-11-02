package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Close() {
	d.watch.Close()

	if c, okay := d.source.(face.Closer); okay {
		c.Close()
	}
}
