package openbao_tester

import (
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/openbao/unit/openbao_server"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	s *openbao_server.Server,
	jwt string,
) *openbao.Client {
	t.Helper()
	path := filepath.Join(t.TempDir(), "token")
	system.SaveFile(path, jwt)
	result := openbao.New(locator.NewOrigin(s.Address())).WithKubernetes(
		"alfa-role",
		path,
	)
	t.Cleanup(result.Close)

	return result
}
