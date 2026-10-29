package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
	"testing"
)

func New(
	t *testing.T,
	s *mock_source.Source,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Owner = "alfa"
	o.Source = func() face.RunSource { return s }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
