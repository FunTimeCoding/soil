package base

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"strconv"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.LitePath = s.LitePath
	o.Project = strconv.Itoa(constant.FixtureProject)
	o.Branch = constant.PublishBranch
	o.AuthorityDirectory = constant.FixtureAuthorityDirectory
	o.SecretAuthority = constant.FixtureClusterAuthority
	o.SecretPath = constant.FixtureSecretPath
	o.Forge = func() face.GitLabClient { return s.Forge }
	o.Clock = s.clock
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
