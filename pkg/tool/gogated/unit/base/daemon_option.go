package base

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
)

func (s *Server) daemonOption() *option.Option {
	result := option.New()
	result.LitePath = s.litePath
	result.Secret = constant.FixtureSecret
	result.Issuer = constant.FixtureIssuer
	result.SuperUserMail = constant.FixtureMail
	result.SuperUserPassword = constant.FixturePassword
	result.AdminClientIdentifier = constant.FixtureAdminClient
	result.AdminClientSecret = constant.FixtureAdminSecret
	result.Directory = func() face.DirectoryClient { return s.directory }

	return result
}
