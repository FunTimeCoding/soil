package unit

import (
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/openbao/unit/upstream_tester"
	"github.com/funtimecoding/soil/pkg/system"
	web "github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"path/filepath"
	"testing"
)

func client(
	t *testing.T,
	u *upstream_tester.Tester,
	jwt string,
) *openbao.Client {
	t.Helper()
	path := filepath.Join(t.TempDir(), "token")
	system.SaveFile(path, jwt)

	return openbao.New(web.Locator(t, u.Address())).WithKubernetes(
		"alfa-role",
		path,
	)
}
