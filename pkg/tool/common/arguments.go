package common

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Arguments(a *argument.Instance) {
	a.String(
		argument.Host,
		environment.Optional(constant.QualifiedName),
		fmt.Sprintf("Host, fallback: %s", constant.QualifiedName),
	)
	a.String(
		argument.Token,
		environment.Fallback(
			TokenEnvironment,
			environment.Optional(constant.JobToken),
		),
		fmt.Sprintf(
			"Token, fallbacks: %s, %s",
			TokenEnvironment,
			constant.JobToken,
		),
	)
	a.String(
		argument.Owner,
		environment.Fallback(
			OwnerEnvironment,
			environment.Optional(constant.ProjectNamespace),
		),
		fmt.Sprintf(
			"Owner, fallbacks: %s, %s",
			OwnerEnvironment,
			constant.ProjectNamespace,
		),
	)
	a.String(
		argument.Repository,
		environment.Optional(RepositoryEnvironment),
		fmt.Sprintf("Repository, fallback: %s", RepositoryEnvironment),
	)
}
