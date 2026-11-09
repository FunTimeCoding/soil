package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
)

type Tester struct {
	Catalog   *catalog.Service
	Query     *query.Service
	Backup    *backup.Service
	Dump      *mock_dump.Program
	Directory string
}
