package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
)

func New(s *mock_source.Source) *service.Service {
	return service.New(s, "alfa")
}
