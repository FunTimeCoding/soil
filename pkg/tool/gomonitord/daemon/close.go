package daemon

import "github.com/funtimecoding/soil/pkg/relational"

func (d *Daemon) Close() {
	relational.CloseMapper(d.database)
}
