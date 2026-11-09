package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump"
	"time"
)

func New() *Postgres {
	return &Postgres{Client: BuildClient, Dump: dump.New(), Clock: time.Now}
}
