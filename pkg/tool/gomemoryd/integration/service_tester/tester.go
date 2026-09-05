package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
)

type Tester struct {
	Service *service.Service
	Indexer *mock_indexer.Indexer
}
