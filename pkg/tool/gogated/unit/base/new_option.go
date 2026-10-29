package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"path/filepath"
	"testing"
)

func NewOption(t *testing.T) *option.Option {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gogated.db")
	result.Secret = constant.FixtureSecret
	result.Issuer = constant.FixtureIssuer
	result.SuperUserMail = constant.FixtureMail
	result.SuperUserPassword = constant.FixturePassword
	result.AdminClientIdentifier = constant.FixtureAdminClient
	result.AdminClientSecret = constant.FixtureAdminSecret

	return result
}
