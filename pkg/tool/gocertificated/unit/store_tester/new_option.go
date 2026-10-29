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
	f face.Forge,
) *option.Certificate {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gocertificated.db")
	o.Project = strconv.Itoa(constant.FixtureProject)
	o.Branch = constant.PublishBranch
	o.AuthorityDirectory = constant.FixtureAuthorityDirectory
	o.SecretAuthority = constant.FixtureClusterAuthority
	o.SecretPath = constant.FixtureSecretPath
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://certificate.example.org"
	o.Forge = func() face.Forge { return f }

	return o
}
