package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/service"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/recording_indexer"
)

type Tester struct {
	Service      *service.Service
	Indexer      *recording_indexer.Indexer
	PackageRoot  string
	KeyDirectory string
	SignatureKey string
}
