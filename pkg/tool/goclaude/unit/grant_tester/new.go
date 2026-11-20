package grant_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"path/filepath"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	result := &Tester{
		t:          t,
		Mock:       mock_client.New(t),
		Session:    "grant-session",
		Repository: t.TempDir(),
		Other:      t.TempDir(),
	}

	for _, r := range []string{result.Repository, result.Other} {
		system.MakeDirectory(filepath.Join(r, constant.GitMetadataDirectory))
	}

	result.Mock.Join(result.Session, 1, time.Now())

	return result
}
