package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/service"
	"testing"
)

type Tester struct {
	Output  string
	parser  face.TopStatsParser
	service *service.Service
	t       *testing.T
}
