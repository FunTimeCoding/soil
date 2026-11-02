package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Host = "directory.test"
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://directory.example.org"
	o.Source = func() face.DirectoryClient { return c }
	result := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
