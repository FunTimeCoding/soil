package gosed

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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
	a.String(argumentConstant.Branch, git.MainBranch, "Branch to commit to")
	a.String(argumentConstant.Path, "", "Path in repository")
	var replaces []string
	a.StringSliceVariable(
		&replaces,
		argumentConstant.Replace,
		nil,
		"One or more prefix replaces (Example: 'image: app:=v1.2.3')",
	)
	var actions []string
	a.StringSliceVariable(
		&actions,
		argumentConstant.Action,
		nil,
		"One or more file:prefix=value actions for multi-file commits",
	)
	a.Parse(version, gitHash, buildDate)
	common.ValidateArguments(a)
	o := option.New()
	o.Host = a.GetString(argumentConstant.Host)
	o.Token = a.GetString(argumentConstant.Token)
	o.Owner = a.GetString(argumentConstant.Owner)
	o.Repository = a.GetString(argumentConstant.Repository)
	o.Branch = a.GetString(argumentConstant.Branch)
	o.Path = a.GetString(argumentConstant.Path)
	o.Replaces = sed.Parse(replaces)
	o.RawActions = actions
	o.Message = a.RequiredPositional(0, "MESSAGE")
	sed.Run(o)
}
