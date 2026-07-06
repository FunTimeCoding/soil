package gocommit

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
	a.String(argument.Branch, git.MainBranch, "Branch to commit to")
	a.String(argument.Path, "", "Path in repository")
	a.String(argument.Template, "", "Template file for commit")
	var replaces []string
	a.StringSliceVariable(
		&replaces,
		argument.Replace,
		nil,
		"One or more key-value pairs to replace (Example: FOO=BAR)",
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
	o.Template = a.GetString(argument.Template)
	o.Replace = replaces
	o.Message = a.RequiredPositional(0, "MESSAGE")
	commit.Run(o)
}
