package common

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	gitlab "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common/constant"
)

func Arguments(a *argument.Instance) {
	a.String(
		argumentConstant.Host,
		environment.Optional(gitlab.QualifiedName),
		fmt.Sprintf("Host, fallback: %s", gitlab.QualifiedName),
	)
	a.String(
		argumentConstant.Token,
		environment.Fallback(
			constant.TokenEnvironment,
			environment.Optional(gitlab.JobToken),
		),
		fmt.Sprintf(
			"Token, fallbacks: %s, %s",
			constant.TokenEnvironment,
			gitlab.JobToken,
		),
	)
	a.String(
		argumentConstant.Owner,
		environment.Fallback(
			constant.OwnerEnvironment,
			environment.Optional(gitlab.ProjectNamespace),
		),
		fmt.Sprintf(
			"Owner, fallbacks: %s, %s",
			constant.OwnerEnvironment,
			gitlab.ProjectNamespace,
		),
	)
	a.String(
		argumentConstant.Repository,
		environment.Optional(constant.RepositoryEnvironment),
		fmt.Sprintf("Repository, fallback: %s", constant.RepositoryEnvironment),
	)
}
