package service_tester

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"testing"
)

func NewWith(
	t *testing.T,
	logCachePath string,
	outputPath string,
	p face.Client,
	r soilFace.Reporter,
) *service.Service {
	t.Helper()

	return NewService(t, NewOption(t, logCachePath, outputPath, p), r)
}
