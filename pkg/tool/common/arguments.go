package common

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common/constant"
)

func Arguments(a *argument.Instance) {
	a.String(
		argumentConstant.Host,
		environment.Optional(gitlabConstant.QualifiedName),
		fmt.Sprintf("Host, fallback: %s", gitlabConstant.QualifiedName),
	)
	a.String(
		argumentConstant.Token,
		environment.Fallback(
			constant.TokenEnvironment,
			environment.Optional(gitlabConstant.JobToken),
		),
		fmt.Sprintf(
			"Token, fallbacks: %s, %s",
			constant.TokenEnvironment,
			gitlabConstant.JobToken,
		),
	)
	a.String(
		argumentConstant.Owner,
		environment.Fallback(
			constant.OwnerEnvironment,
			environment.Optional(gitlabConstant.ProjectNamespace),
		),
		fmt.Sprintf(
			"Owner, fallbacks: %s, %s",
			constant.OwnerEnvironment,
			gitlabConstant.ProjectNamespace,
		),
	)
	a.String(
		argumentConstant.Repository,
		environment.Optional(constant.RepositoryEnvironment),
		fmt.Sprintf("Repository, fallback: %s", constant.RepositoryEnvironment),
	)
}
