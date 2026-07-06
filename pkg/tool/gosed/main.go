package gosed

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/gosed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosed/sed"
	"github.com/funtimecoding/soil/pkg/tool/gosed/sed/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	common.Arguments(a)
	a.String(argument.Branch, git.MainBranch, "Branch to commit to")
	a.String(argument.Path, "", "Path in repository")
	var replaces []string
	a.StringSliceVariable(
		&replaces,
		argument.Replace,
		nil,
		"One or more prefix replaces (Example: 'image: app:=v1.2.3')",
	)
	var actions []string
	a.StringSliceVariable(
		&actions,
		argument.Action,
		nil,
		"One or more file:prefix=value actions for multi-file commits",
	)
	a.Parse(version, gitHash, buildDate)
	common.ValidateArguments(a)
	o := option.New()
	o.Host = a.GetString(argument.Host)
	o.Token = a.GetString(argument.Token)
	o.Owner = a.GetString(argument.Owner)
	o.Repository = a.GetString(argument.Repository)
	o.Branch = a.GetString(argument.Branch)
	o.Path = a.GetString(argument.Path)
	o.Replaces = sed.Parse(replaces)
	o.RawActions = actions
	o.Message = a.RequiredPositional(0, "MESSAGE")
	sed.Run(o)
}
