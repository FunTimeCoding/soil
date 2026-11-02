package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/service"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	p face.TopStatsParser,
) (*service.Service, string) {
	t.Helper()
	output := t.TempDir()

	return service.New(
		p,
		filepath.Join(t.TempDir(), "TW5_Top_Stat_Parse.html"),
		output,
		logger.New(t.Context()),
	), output
}
