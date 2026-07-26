package gocommit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/gocommit/commit"
	"github.com/funtimecoding/soil/pkg/tool/gocommit/commit/option"
	"github.com/funtimecoding/soil/pkg/tool/gocommit/constant"
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
	a.String(argumentConstant.Template, "", "Template file for commit")
	var replaces []string
	a.StringSliceVariable(
		&replaces,
		argumentConstant.Replace,
		nil,
		"One or more key-value pairs to replace (Example: FOO=BAR)",
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
	o.Template = a.GetString(argumentConstant.Template)
	o.Replace = replaces
	o.Message = a.RequiredPositional(0, "MESSAGE")
	commit.Run(o)
}
