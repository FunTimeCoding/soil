package option

import (
	"github.com/funtimecoding/soil/pkg/postgres/face"
	gopostgresd "github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Postgres struct {
	option.Web
	Inventory       *instance.Inventory
	Client          func(*instance.Instance) (face.PostgresClient, error)
	BackupDirectory string
	Dump            gopostgresd.Dump
	Clock           func() time.Time
}
