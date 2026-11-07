package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (d *Daemon) Close() {
	d.guard.Close()

	if c, okay := d.forge.(face.Closer); okay {
		c.Close()
	}

	relational.CloseMapper(d.database)
}
