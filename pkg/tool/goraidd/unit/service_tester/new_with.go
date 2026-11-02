package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	parser "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"testing"
)

func NewWith(
	t *testing.T,
	logCachePath string,
	outputPath string,
	p parser.Client,
	r face.Reporter,
) *service.Service {
	t.Helper()

	return NewService(t, NewOption(t, logCachePath, outputPath, p), r)
}
