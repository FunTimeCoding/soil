package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/postgres/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
	"testing"
)

type Tester struct {
	instances map[string]*mock_client.Client
	Reporter  *memory.Memory
	Dump      *mock_dump.Program
	Directory string
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	catalog   *catalog.Service
	query     *query.Service
	backup    *backup.Service
	t         *testing.T
}
