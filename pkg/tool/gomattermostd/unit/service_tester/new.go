package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	m := mock_client.New(t, "kilo")
	result := &Tester{
		Mattermost: m,
		Reporter:   memory.New(),
		Path:       filepath.Join(t.TempDir(), "mattermost.db"),
		client:     m,
		t:          t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
