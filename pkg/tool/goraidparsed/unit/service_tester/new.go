package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/service"
	"testing"
)

func New(t *testing.T) *service.Service {
	t.Helper()
	o := option.New()
	o.ParserPath = t.TempDir()
	o.TemplatePath = t.TempDir()
	o.OutputPath = t.TempDir()

	return service.New(o, logger.New(t.Context()))
}
