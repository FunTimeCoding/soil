package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
)

type Server struct {
	catalog   *catalog.Service
	query     *query.Service
	backup    *backup.Service
	inventory *instance.Inventory
	reporter  face.Reporter
}
