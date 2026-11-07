package store_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"path/filepath"
	"strconv"
	"testing"
)

func NewOption(
	t *testing.T,
	f face.GitLabClient,
) *option.Certificate {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gocertificated.db")
	o.Project = strconv.Itoa(constant.FixtureProject)
	o.Branch = constant.PublishBranch
	o.AuthorityDirectory = constant.FixtureAuthorityDirectory
	o.SecretAuthority = constant.FixtureClusterAuthority
	o.SecretPath = constant.FixtureSecretPath
	o.Forge = func() face.GitLabClient { return f }

	return o
}
