package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"path/filepath"
	"testing"
)

func newDaemonOption(
	t *testing.T,
) (*option.Option, *memoryMock.Client, *queryMock.Client) {
	t.Helper()
	c := mock_client.New()
	c.Harbor = t.TempDir()
	m := memoryMock.New()
	q := queryMock.New()
	n := mock_notifier.New()
	sources := []string{t.TempDir()}
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "goclauded.sqlite")
	result.SessionExportPath = t.TempDir()
	result.Harbor = c.Harbor
	result.Claude = func() face.ClaudeSource { return c }
	result.Memory = func() memoryd.Client { return m }
	result.Query = func() queryd.Client { return q }
	result.Notifier = func() face.Notifier { return n }
	result.Sources = func() []string { return sources }

	return result, m, q
}
