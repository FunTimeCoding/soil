package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	logCachePath string,
	outputPath string,
	p face.Client,
) *option.Raid {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "goraidd.db")
	result.LogCachePath = logCachePath
	result.ElitePath = t.TempDir()
	result.OutputPath = outputPath
	result.Parser = func() face.Client { return p }

	return result
}
