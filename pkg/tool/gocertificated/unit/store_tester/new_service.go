package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func NewService(s *store.Store) *service.Service {
	return service.New(s, nil)
}
